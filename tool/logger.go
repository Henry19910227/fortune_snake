package tool

import (
	"bytes"
	"encoding/json"
	"fmt"
	model "game_server_slots_fortune_snake/internal/model/config/system"
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	instance Logger
	once     sync.Once
)

// Log 全局日志实例
//var Log = &Logger{}

// LogConfig 结构体，存储日志相关配置
type LogConfig struct {
	LogDir        string
	ElkEnabled    bool
	ElkIndex      string
	BatchSize     int
	BatchInterval time.Duration
}

// Logger 结构体，存储本地日志和 ELK 日志队列
type logger struct {
	LocalLogger *zap.Logger
	ElkClient   *elasticsearch.Client
	Config      LogConfig
	elkLogQueue chan map[string]interface{}
	wg          sync.WaitGroup
	bufferPool  *sync.Pool
	closeOnce   sync.Once
}

// InitLogger 初始化全局日志实例
//func InitLogger(elkClient *elasticsearch.Client) {
//	Log = NewLogger(elkClient)
//}

// NewLogger 创建 Logger 实例
func NewLogger(config *model.Config) Logger {
	// 使用 `filepath.Join` 确保路径拼接正确
	logFilePath := filepath.Join(config.Log.LogDir, fmt.Sprintf("app_%s.log", time.Now().Format("2006-01-02")))
	// 本地日志切割配置
	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    config.Log.MaxSize,
		MaxBackups: config.Log.MaxBackups,
		MaxAge:     config.Log.MaxAge,
		Compress:   config.Log.Compress,
	})
	// 设置日志编码格式
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	// 创建日志核心
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout)), zapcore.InfoLevel)
	localLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	l := &logger{
		LocalLogger: localLogger,
		Config: LogConfig{
			LogDir:        config.Log.LogDir,                  // LogDir 是本地日志文件的保存路
			ElkEnabled:    config.Elasticsearch.Enabled,       // ElkEnabled 控制是否开启 ELK 日志写入（异步方式）
			ElkIndex:      config.Elasticsearch.IndexPrefix,   // ElkIndex 是写入 Elasticsearch 的索引前缀
			BatchSize:     config.Elasticsearch.BatchSize,     // BatchSize 是写入 ELK 的最大日志条数阈值，达到该值即触发批量发送
			BatchInterval: config.Elasticsearch.BatchInterval, // BatchInterval 是日志批量写入 ELK 的时间间隔
		},
		elkLogQueue: make(chan map[string]interface{}, config.Elasticsearch.BatchSize*2),
		bufferPool:  &sync.Pool{New: func() interface{} { return new(bytes.Buffer) }},
	}
	// 启动 ELK 日志处理协程
	if l.Config.ElkEnabled && l.ElkClient != nil {
		l.wg.Add(1)
		go l.processElkLogs()
	}
	return l
}

// Logs 记录日志
func (l *logger) Logs(level string, writeToElk bool, msg string, fields ...zap.Field) {
	switch level {
	case "info":
		l.LocalLogger.Info(msg, fields...)
	case "error":
		l.LocalLogger.Error(msg, fields...)
	case "debug":
		l.LocalLogger.Debug(msg, fields...)
	case "warn":
		l.LocalLogger.Warn(msg, fields...)
	default:
		l.LocalLogger.Warn("未知日志级别", zap.String("level", level), zap.String("message", msg))
	}
	// 仅在 `writeToElk` 为 true 时写入 ELK
	if writeToElk {
		l.queueElkLog(level, msg, fields...)
	}
}

func (l *logger) SetELK(elkClient *elasticsearch.Client) {
	l.ElkClient = elkClient
	// 启动 ELK 日志处理协程
	if l.Config.ElkEnabled && l.ElkClient != nil {
		l.wg.Add(1)
		go l.processElkLogs()
	}
}

// queueElkLog 将日志添加到 ELK 队列
func (l *logger) queueElkLog(level, msg string, fields ...zap.Field) {
	if !l.Config.ElkEnabled || l.ElkClient == nil {
		return
	}
	// 过滤 nil 值，防止 `json.Marshal` 失败
	fieldMap := make(map[string]interface{})
	for _, field := range fields {
		if field.Interface != nil {
			fieldMap[field.Key] = field.Interface
		}
	}
	logData := map[string]interface{}{
		"timestamp": time.Now().Format(time.RFC3339),
		"level":     level,
		"message":   msg,
		"fields":    fieldMap,
	}
	// 防止队列阻塞
	select {
	case l.elkLogQueue <- logData:
	default:
		log.Println("⚠️ ELK 日志队列已满，日志可能丢失！")
	}
}

// processElkLogs 批量处理 ELK 日志
func (l *logger) processElkLogs() {
	defer l.wg.Done()
	ticker := time.NewTicker(l.Config.BatchInterval)
	defer ticker.Stop()
	logs := make([]map[string]interface{}, 0, l.Config.BatchSize)
	for {
		select {
		case logData, ok := <-l.elkLogQueue:
			if !ok {
				if len(logs) > 0 {
					l.bulkSendToElk(logs)
				}
				return
			}
			logs = append(logs, logData)
			if len(logs) >= l.Config.BatchSize {
				l.bulkSendToElk(logs)
				logs = logs[:0]
			}
		case <-ticker.C:
			if len(logs) > 0 {
				l.bulkSendToElk(logs)
				logs = logs[:0]
			}
		}
	}
}

// bulkSendToElk 批量提交日志到 ELK
func (l *logger) bulkSendToElk(logs []map[string]interface{}) {
	if len(logs) == 0 || l.ElkClient == nil {
		return
	}
	indexName := fmt.Sprintf("%s-%s", l.Config.ElkIndex, time.Now().Format("2006-01-02"))
	buffer := l.bufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer l.bufferPool.Put(buffer)
	for _, logData := range logs {
		meta := map[string]interface{}{
			"index": map[string]interface{}{
				"_index": indexName,
			},
		}
		metaJSON, _ := json.Marshal(meta)
		logJSON, _ := json.Marshal(logData)

		buffer.Write(metaJSON)
		buffer.WriteString("\n")
		buffer.Write(logJSON)
		buffer.WriteString("\n")
	}
	res, err := l.ElkClient.Bulk(bytes.NewReader(buffer.Bytes()))
	if err != nil || res.IsError() {
		log.Printf("❌ ELK 批量写入失败: %v", err)
	}
}

// Close 关闭 Logger
func (l *logger) Close() {
	l.closeOnce.Do(func() {
		close(l.elkLogQueue)
	})
	l.wg.Wait()
	// 安全处理 `Sync()` 方法
	if err := l.LocalLogger.Sync(); err != nil && err.Error() != "invalid argument" {
		log.Printf("⚠️ 日志同步失败: %v", err)
	}
}

func InitLoggerInstance(config *model.Config) {
	once.Do(func() {
		instance = NewLogger(config)
	})
}

func Log() Logger {
	return instance
}
