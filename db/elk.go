package db

import (
	"crypto/tls"
	"fmt"
	"game_server_slots_fortune_snake/internal/model"
	"game_server_slots_fortune_snake/tool"
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type elkDB struct {
	client *elasticsearch.Client
}

func NewElkDB(config *model.Config) ElkDB {
	esURL := fmt.Sprintf("%s:%d", config.Elasticsearch.Host, config.Elasticsearch.Port)
	// 处理 HTTPS 连接
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,             // 是否跳过 TLS 证书验证（生产环境设为 false）
			MinVersion:         tls.VersionTLS12, // 设置最小 TLS 版本为 1.2
		},
		MaxIdleConns:          config.Elasticsearch.MaxIdleConns,        // 最大空闲连接数
		MaxConnsPerHost:       config.Elasticsearch.MaxConnsPerHost,     // 每个主机的最大连接数
		MaxIdleConnsPerHost:   config.Elasticsearch.MaxIdleConnsPerHost, // 每个主机的最大空闲连接数
		IdleConnTimeout:       config.Elasticsearch.KeepAlive,           // 连接的 Keep-Alive 时间
		ResponseHeaderTimeout: config.Elasticsearch.RequestTimeout,      // 等待响应头的超时时间
	}
	cfg := elasticsearch.Config{
		Addresses:     []string{esURL},                 // Elasticsearch 服务器地址
		Username:      config.Elasticsearch.Username,   // 用户名
		Password:      config.Elasticsearch.Password,   // 密码
		Transport:     tr,                              // 自定义 HTTP 传输配置
		RetryOnStatus: []int{502, 503, 504},            // 在 502/503/504 状态码时重试
		MaxRetries:    config.Elasticsearch.MaxRetries, // 最大重试次数
	}
	var client *elasticsearch.Client
	var err error
	// 尝试连接到 Elasticsearch，最多重试 MaxRetries 次
	for i := 0; i <= config.Elasticsearch.MaxRetries; i++ {
		client, err = elasticsearch.NewClient(cfg)
		if err == nil {
			break // 成功连接后跳出循环
		}
		tool.Log().Logs("error", true, fmt.Sprintf("⚠️ [ELK ERROR] Connection attempt %d failed", i+1), zap.String("Component", "Elasticsearch"), zap.Error(err))
		time.Sleep(config.Elasticsearch.RetryWaitTime) // 等待一段时间后重试
	}
	if err != nil {
		tool.Log().Logs("error", true, fmt.Sprintf("❌ [ELK ERROR] Failed to connect to Elasticsearch after %d attempts", config.Elasticsearch.MaxRetries), zap.String("Component", "Elasticsearch"), zap.Error(err))
		return &elkDB{client: nil}
	}
	tool.Log().Logs("info", false, fmt.Sprintf("✅ [ELK] Successfully connected to Elasticsearch at %s", esURL), zap.String("Component", "Elasticsearch"))
	return &elkDB{client: client}
}

func (e elkDB) Client() *elasticsearch.Client {
	return e.client
}
