package model

type Config struct {
	Server   ServerConfig   `yaml:"app"`
	Language LanguageConfig `yaml:"language"`
	Pprof    PprofConfig    `yaml:"pprof"`
	Etcd     EtcdConfig     `yaml:"etcd"`
	Grpc     GrpcConfig     `yaml:"grpc_server"`
	Log      LogConfig      `yaml:"log"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
	Mongodb  MongodbConfig  `yaml:"mongodb"`
	Kafka    KafkaConfig    `yaml:"kafka"`
}

// ServerConfig 对应 server 部分
type ServerConfig struct {
	AppName    string `yaml:"app_name"`    // 服务名称
	Ip         string `yaml:"ip"`          // 服务监听IP(rpc服务访问使用)
	Port       int    `yaml:"port"`        // 服务监听端口
	ServerNode int64  `yaml:"server_node"` // 服务器节点
}

// LanguageConfig 对应 language 部分
type LanguageConfig struct {
	Default          string   `yaml:"default"`           // 默认语言包
	SupportLanguages []string `yaml:"support_languages"` // 支持的语言列表
}

// PprofConfig 对应 pprof 部分
type PprofConfig struct {
	Enabled bool `yaml:"enabled"` // 是否启用 pprof
	Port    int  `yaml:"port"`    // pprof 分析服务器端口
}

// EtcdConfig 对应 etcd 部分
type EtcdConfig struct {
	Endpoints     []string `yaml:"endpoints"`      // Etcd 地址列表
	ServicePrefix string   `yaml:"service_prefix"` // 服务前缀
	Ttl           int      `yaml:"ttl"`            // 注册信息的过期时间（秒）
}

// GrpcConfig 对应 grpc_server 部分
type GrpcConfig struct {
	UseRandom   bool                      `yaml:"use_random"`   // 是否随机选择 RPC（true 随机，false 轮询）
	MaxRetries  int                       `yaml:"max_retries"`  // 最大重试次数
	RetryDelay  int                       `yaml:"retry_delay"`  // 重试间隔（秒）
	EtcdConfigs map[string]GrpcEtcdConfig `yaml:"etcd_configs"` // 不同服务的 etcd 配置
}

// GrpcEtcdConfig 对应 grpc_server.etcd_configs 下的配置项
type GrpcEtcdConfig struct {
	Endpoints     []string `yaml:"endpoints"`      // 服务对应的 etcd 地址
	ServicePrefix string   `yaml:"service_prefix"` // 服务前缀
}

// JwtConfig 对应 jwt 部分
type JwtConfig struct {
	Secret                  string `yaml:"secret"`                     // JWT 密钥
	JwtTtl                  int    `yaml:"jwt_ttl"`                    // Token 有效期（秒）
	JwtBlacklistGracePeriod int    `yaml:"jwt_blacklist_grace_period"` // 黑名单宽限时间（秒）
	RefreshGracePeriod      int    `yaml:"refresh_grace_period"`       // Token 自动刷新宽限时间（秒）
}

// LogConfig 对应 log 部分
type LogConfig struct {
	Level      string `yaml:"level"`       // 日志等级
	RootDir    string `yaml:"root_dir"`    // 日志目录
	Filename   string `yaml:"filename"`    // 日志文件名称
	Format     string `yaml:"format"`      // 日志格式（如 json）
	ShowLine   bool   `yaml:"show_line"`   // 是否显示调用行
	MaxBackups int    `yaml:"max_backups"` // 旧文件的最大个数
	MaxSize    int    `yaml:"max_size"`    // 单个日志文件最大大小（MB）
	MaxAge     int    `yaml:"max_age"`     // 旧日志最大保留天数
	Compress   bool   `yaml:"compress"`    // 是否压缩日志
}

// DatabaseConfig 对应 database 部分
type DatabaseConfig struct {
	Driver              string `yaml:"driver"`                 // 数据库驱动
	Host                string `yaml:"host"`                   // 数据库地址
	Port                int    `yaml:"port"`                   // 数据库端口
	Database            string `yaml:"database"`               // 数据库名称
	Username            string `yaml:"username"`               // 数据库用户名
	Password            string `yaml:"password"`               // 数据库密码
	Charset             string `yaml:"charset"`                // 编码格式
	MaxIdleConns        int    `yaml:"max_idle_conns"`         // 空闲连接数
	MaxOpenConns        int    `yaml:"max_open_conns"`         // 最大连接数
	LogMode             string `yaml:"log_mode"`               // 日志级别
	EnableFileLogWriter bool   `yaml:"enable_file_log_writer"` // 是否启用文件日志
	LogFilename         string `yaml:"log_filename"`           // SQL 日志文件名称
}

// RedisConfig 对应 redis 部分
type RedisConfig struct {
	MasterName       string   `yaml:"master_name"`       // Redis 主节点名称
	SentinelAddress  []string `yaml:"sentinel_address"`  // Sentinel 地址列表
	Password         string   `yaml:"password"`          // Redis 主/从密码
	SentinelPassword string   `yaml:"sentinel_password"` // Redis Sentinel 密码
	Db               int      `yaml:"db"`                // Redis 数据库索引
}

// MongodbConfig 对应 mongodb 部分
type MongodbConfig struct {
	Username   string   `yaml:"username"`    // MongoDB 用户名
	Password   string   `yaml:"password"`    // MongoDB 密码
	Hosts      []string `yaml:"hosts"`       // MongoDB 地址列表
	Database   string   `yaml:"database"`    // MongoDB 数据库名称
	AuthSource string   `yaml:"auth_source"` // 认证数据库
}

// KafkaConfig 结构体用于解析 YAML 配置
type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`  // Kafka broker 列表（支持多个）
	GroupID string   `yaml:"group_id"` // Kafka 消费者组 ID
	Version string   `yaml:"version"`  // Kafka 服务器版本
	Oldest  bool     `yaml:"oldest"`   // 是否从最早的 offset 开始消费（true: earliest, false: latest）
	// SASL 认证配置（适用于 Kafka 认证）
	SASL struct {
		Enabled  bool   `yaml:"enabled"`  // 是否启用 SASL 认证
		User     string `yaml:"user"`     // SASL 用户名
		Password string `yaml:"password"` // SASL 密码
	} `yaml:"sasl"`
	// TLS 加密配置（适用于 Kafka 安全连接）
	TLS struct {
		Enabled    bool   `yaml:"enabled"`      // 是否启用 TLS 加密
		CACertFile string `yaml:"ca_cert_file"` // TLS CA 证书文件路径（可选）
	} `yaml:"tls"`
	// 生产者配置（用于优化 Kafka 生产端）
	Producer struct {
		Acks        string `yaml:"acks"`        // 生产者消息确认模式 ("all", "1", "0")
		Retries     int    `yaml:"retries"`     // 生产者最大重试次数
		Compression string `yaml:"compression"` // 消息压缩类型（gzip, snappy, lz4, zstd）
		BatchSize   int    `yaml:"batch_size"`  // 生产者批量发送的消息大小（单位: 字节）
		LingerMs    int    `yaml:"linger_ms"`   // 生产者等待时间（单位: 毫秒）
	} `yaml:"producer"`
	// 消费者配置（用于优化 Kafka 消费端）
	Consumer struct {
		AutoCommit           bool `yaml:"auto_commit"`             // 是否自动提交 offset
		AutoCommitIntervalMs int  `yaml:"auto_commit_interval_ms"` // 自动提交 offset 的时间间隔（毫秒）
		FetchMinBytes        int  `yaml:"fetch_min_bytes"`         // 消费者最小拉取消息大小
		FetchMaxBytes        int  `yaml:"fetch_max_bytes"`         // 消费者最大拉取消息大小
		SessionTimeoutMs     int  `yaml:"session_timeout_ms"`      // 消费者 session 超时时间（毫秒）
	} `yaml:"consumer"`
	// 游戏 Kafka 主题映射
	Games map[string]struct {
		RobotBet   string `yaml:"robot_bet"`   // 下注事件 Kafka 主题
		GameResult string `yaml:"game_result"` // 游戏结果 Kafka 主题
	} `yaml:"games"`
}

// RateLimitConfig 对应 rate-limit 部分
type RateLimitConfig struct {
	FillInterval int `yaml:"fill_interval"` // 令牌桶填充时间间隔（秒）
	Capacity     int `yaml:"capacity"`      // 令牌桶容量
}
