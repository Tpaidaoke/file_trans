package database

import (
	"context"
	"daoke.com/file_trans/conf"
	"daoke.com/file_trans/logger"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBClient *gorm.DB
var RClient *redis.Client

// InitDB 初始化数据库
func InitDB() {
	mysqlConf := conf.AppConfig.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.DBName)
	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default,
	})

	if err != nil {
		logger.Error("数据库连接失败：%v", err.Error())
	}

	// 获取底层 sql.DB 实例，设置连接池参数
	sqlDB, err := client.DB()
	if err != nil {
		logger.Error("获取数据库连接池失败：%v", err.Error())
	}

	sqlDB.SetMaxOpenConns(mysqlConf.MaxOpenConns)       // 最大打开连接数
	sqlDB.SetMaxIdleConns(mysqlConf.MaxIdleConns)       // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(mysqlConf.ConnMaxLifetime) // 连接最大存活时间

	DBClient = client
	logger.Info("连接数据库成功")
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return DBClient
}

// CloseDB 关闭数据库连接
func CloseDB() {
	if DBClient != nil {
		sqlDB, err := DBClient.DB()
		if err != nil {
			logger.Error("关闭数据库连接失败：%v", err.Error())
		}
		// 关闭数据库连接
		if err = sqlDB.Close(); err != nil {
			return
		}
	}
}

// InitRedis 初始化Redis
func InitRedis() {
	redisConf := conf.AppConfig.Redis
	rClient := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port),
		Password:     redisConf.Password,
		DB:           redisConf.DB,
		PoolSize:     redisConf.PoolSize,     // 连接池大小
		MinIdleConns: redisConf.MinIdleConns, // 最小空闲连接数
		DialTimeout:  redisConf.DialTimeout,  // 连接建立超时时间
	})

	RClient = rClient
	str, err := RClient.Ping(context.Background()).Result()
	if err != nil {
		logger.Error("连接Redis失败：%v", err.Error())
	}
	logger.Info("连接Redis成功", str)
}
