package conf

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

// App 应用基础配置
type App struct {
	Name string `yaml:"name"` // 应用名称
	Env  string `yaml:"env"`  // 运行环境（dev/test/prod）
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port         int           `yaml:"port"`          // 服务监听端口
	ReadTimeout  time.Duration `yaml:"read_timeout"`  // HTTP读取超时时间
	WriteTimeout time.Duration `yaml:"write_timeout"` // HTTP写入超时时间
	IdleTimeout  time.Duration `yaml:"idle_timeout"`  // 连接空闲超时时间
}

// DBConfig 数据库配置
type DBConfig struct {
	Host            string        `yaml:"host"`              // 数据库主机地址
	Port            int           `yaml:"port"`              // 数据库端口
	User            string        `yaml:"user"`              // 数据库用户名
	Password        string        `yaml:"password"`          // 数据库密码
	DBName          string        `yaml:"dbname"`            // 数据库名称
	Charset         string        `yaml:"charset"`           // 字符集
	MaxOpenConns    int           `yaml:"max_open_conns"`    // 最大打开连接数
	MaxIdleConns    int           `yaml:"max_idle_conns"`    // 最大空闲连接数
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"` // 连接最大存活时间
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host         string        `yaml:"host"`           // Redis主机地址
	Port         int           `yaml:"port"`           // Redis端口
	Password     string        `yaml:"password"`       // Redis密码
	DB           int           `yaml:"db"`             // Redis数据库编号
	PoolSize     int           `yaml:"pool_size"`      // 连接池大小
	MinIdleConns int           `yaml:"min_idle_conns"` // 最小空闲连接数
	DialTimeout  time.Duration `yaml:"dial_timeout"`   // 连接建立超时时间
}

// MinIOConfig MinIO配置
type MinIOConfig struct {
	Endpoint   string `yaml:"endpoint"`    // 服务地址（host:port）
	AccessKey  string `yaml:"access_key"`  // 访问密钥（AK）
	SecretKey  string `yaml:"secret_key"`  // 密钥（SK）
	UseSSL     bool   `yaml:"use_ssl"`     // 是否启用SSL
	BucketName string `yaml:"bucket_name"` // 存储桶名称
	MaxRetries int    `yaml:"max_retries"` // 最大重试次数
}

// LogFileConfig 日志文件配置（嵌套结构体）
type LogFileConfig struct {
	Path      string `yaml:"path"`       // 日志文件存储路径
	MaxSize   int    `yaml:"max_size"`   // 单个文件最大大小（MB）
	MaxBackup int    `yaml:"max_backup"` // 最大备份文件数
	MaxAge    int    `yaml:"max_age"`    // 日志保留天数
	Compress  bool   `yaml:"compress"`   // 是否压缩备份文件
}

// LogConfig 日志配置
type LogConfig struct {
	Level  string        `yaml:"level"`  // 日志级别（debug/info/warn/error）
	Output string        `yaml:"output"` // 输出方式（console/file/both）
	File   LogFileConfig `yaml:"file"`   // 日志文件配置（嵌套）
}

// Config 聚合所有配置
type Config struct {
	App    App          `yaml:"app"`    // 应用基础配置
	Server ServerConfig `yaml:"server"` // 服务器配置
	DB     DBConfig     `yaml:"db"`     // 数据库配置
	Redis  RedisConfig  `yaml:"redis"`  // Redis配置
	MinIO  MinIOConfig  `yaml:"minio"`  // MinIO配置
	Log    LogConfig    `yaml:"log"`    // 日志配置
}

var AppConfig Config // 全局配置变量

func InitConfig(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("读取配置文件失败：%v", err)
	}

	if err = yaml.Unmarshal(data, &AppConfig); err != nil {
		log.Fatalf("解析配置文件失败：%v", err)
	}

	log.Println("配置文件加载成功")
}
