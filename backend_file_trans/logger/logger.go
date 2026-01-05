// logger/logger.go
package logger

import (
	"context"
	"daoke.com/file_trans/conf"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"gorm.io/gorm/logger" // 导入GORM日志接口
)

var globalLogger *slog.Logger // 全局日志实例
var Default logger.Interface  // 导出的Default，实现GORM的Logger接口

// 自定义GORM日志适配器，实现gorm/logger.Interface
type gormLogger struct {
	level logger.LogLevel // GORM日志级别
}

// InitLogger 初始化日志（同时初始化GORM的Default日志）
func InitLogger() {
	// 1. 解析日志级别
	level := new(slog.Level)
	switch conf.AppConfig.Log.Level {
	case "debug":
		*level = slog.LevelDebug
	case "warn":
		*level = slog.LevelWarn
	case "error":
		*level = slog.LevelError
	default: // 默认info级别
		*level = slog.LevelInfo
	}

	// 2. 配置日志处理器（输出到控制台或文件）
	var handler slog.Handler
	switch conf.AppConfig.Log.Output {
	case "file":
		// 确保日志目录存在
		if err := os.MkdirAll(conf.AppConfig.Log.File.Path, 0755); err != nil {
			panic("创建日志目录失败: " + err.Error())
		}
		// 日志文件名（按日期区分，如 2024-10-01.log）
		fileName := time.Now().Format("2006-01-02") + ".log"
		logFile, err := os.OpenFile(
			filepath.Join(conf.AppConfig.Log.File.Path, fileName),
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0644,
		)
		if err != nil {
			panic("打开日志文件失败: " + err.Error())
		}
		// 输出到文件（JSON格式）
		handler = slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: level})
	default: // 输出到控制台（开发环境友好的文本格式）
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	}

	// 3. 初始化全局日志实例
	globalLogger = slog.New(handler).With(
		"app", conf.AppConfig.App.Name,
		"env", conf.AppConfig.App.Env,
	)
	globalLogger.Info("日志模块初始化成功", "level", conf.AppConfig.Log.Level, "output", conf.AppConfig.Log.Output)

	// 4. 初始化GORM的Default日志适配器
	var gormLevel logger.LogLevel
	switch *level {
	case slog.LevelDebug:
		gormLevel = logger.Info // GORM的Info级别包含debug信息
	case slog.LevelInfo:
		gormLevel = logger.Info
	case slog.LevelWarn:
		gormLevel = logger.Warn
	case slog.LevelError:
		gormLevel = logger.Error
	default:
		gormLevel = logger.Info
	}
	Default = &gormLogger{level: gormLevel}
}

// ------------------------------
// 实现gorm/logger.Interface接口方法
// ------------------------------

// LogMode 设置GORM日志级别
func (l *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return &gormLogger{level: level}
}

// Info 输出Info级别日志（GORM的普通信息）
func (l *gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Info {
		Info(msg, data...)
	}
}

// Warn 输出Warn级别日志（GORM的警告信息）
func (l *gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Warn {
		Warn(msg, data...)
	}
}

// Error 输出Error级别日志（GORM的错误信息）
func (l *gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Error {
		Error(msg, data...)
	}
}

// Trace 输出SQL执行轨迹（GORM的SQL日志）
func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.level <= logger.Silent {
		return
	}

	// 计算SQL执行时间
	elapsed := time.Since(begin)
	// 获取SQL语句和影响行数
	sql, rows := fc()

	// 根据错误类型输出不同级别日志
	if err != nil {
		Error("SQL执行错误", "error", err, "sql", sql, "elapsed", elapsed, "rows", rows)
	} else if l.level >= logger.Info {
		Info("SQL执行信息", "sql", sql, "elapsed", elapsed, "rows", rows)
	}
}

// Debug 输出Debug级别日志（GORM的调试信息）
func Debug(msg string, args ...any) {
	globalLogger.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	globalLogger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	globalLogger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	globalLogger.Error(msg, args...)
}
