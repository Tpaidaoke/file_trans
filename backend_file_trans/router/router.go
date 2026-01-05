package router

import (
	"daoke.com/file_trans/controller"
	"daoke.com/file_trans/logger"
	"github.com/gin-gonic/gin"
	"time"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(func(ctx *gin.Context) {
		//中间件
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Origin,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		// 预处理请求
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	})
	// 创建发送控制器
	sendController := controller.NewSendController()
	receiveController := controller.NewReceiveController()

	v1 := r.Group("api/v1")
	{
		// 发送文件
		v1.POST("/sendPackage", sendController.Send)

		// 发送记录
		v1.GET("/sendRecords", sendController.QuerySendRecords)

		// 下载文件
		v1.GET("/receivePackage", receiveController.Receive)

		// 更新文件状态
		v1.POST("/updateFileStatus", receiveController.UpdateFileStatus)

		// 取件记录
		v1.GET("/receiveRecords", receiveController.QueryReceiveRecords)
	}
	return r
}

// RequestLogMiddleware 记录HTTP请求日志
func RequestLogMiddleware(ctx *gin.Context) {
	// 开始时间
	start := time.Now()

	// 处理请求
	ctx.Next()

	// 计算耗时
	cost := time.Since(start)

	// 记录日志（包含：方法、路径、状态码、耗时、客户端IP）
	logger.Info(
		"请求访问记录",
		"method", ctx.Request.Method,
		"path", ctx.Request.URL.Path,
		"status", ctx.Writer.Status(),
		"cost", cost,
		"client_ip", ctx.ClientIP(),
	)
}
