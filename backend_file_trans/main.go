package main

import (
	"daoke.com/file_trans/conf"
	"daoke.com/file_trans/database"
	"daoke.com/file_trans/logger"
	"daoke.com/file_trans/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 初始化配置
	conf.InitConfig("./config.yaml")

	// 根据环境设置Gin模式（新增代码）
	if conf.AppConfig.App.Env == "prod" {
		gin.SetMode(gin.ReleaseMode) // 生产环境切换到release模式
	} else {
		gin.SetMode(gin.DebugMode) // 开发/测试环境使用debug模式
	}

	// 初始化日志
	logger.InitLogger()
	// 初始化数据库
	database.InitDB()
	// 初始化Redis
	database.InitRedis()
	// 初始化MinIO
	database.InitMinIO()
	// 初始化路由
	r := router.InitRouter()
	// 启动服务器
	if err := r.Run(fmt.Sprintf(":%d", conf.AppConfig.Server.Port)); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
