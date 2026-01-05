package controller

import (
	"daoke.com/file_trans/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

type ReceiveController struct {
	ReceiveService *service.ReceiveService
}

func NewReceiveController() *ReceiveController {
	return &ReceiveController{
		ReceiveService: service.NewReceiveService(),
	}
}

func (r *ReceiveController) Receive(ctx *gin.Context) {
	// 从请求参数中获取取件码
	pickupCode := ctx.Query("pickupCode")
	if pickupCode == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请输入取件码！"})
		return
	}

	// 验证取件码格式（6位数字）
	match, _ := regexp.MatchString(`^\d{6}$`, pickupCode)
	if !match {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请输入正确取件码！"})
		return
	}

	// 调用服务层方法查询文件传输信息
	transInfo, err := r.ReceiveService.GetTransInfoByPickupCode(pickupCode)
	if err != nil {
		// 区分不同错误类型返回对应信息
		switch err.Error() {
		case "pickupCode is expired", "pickupCode not found":
			ctx.JSON(http.StatusNotFound, gin.H{"error": "取件码已过期或不存在！"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	// 返回文件下载链接
	ctx.JSON(http.StatusOK, gin.H{
		"fileName":        transInfo.FileName,
		"fileUuid":        transInfo.FileUuid,
		"fileSize":        transInfo.FileSize,
		"expired":         transInfo.IsExpire,
		"fileDownloadUrl": transInfo.StorageUrl,
	})
}

// UpdateFileStatus 处理文件下载完成后的状态更新
func (r *ReceiveController) UpdateFileStatus(ctx *gin.Context) {
	// 定义接收JSON参数的结构体
	type req struct {
		FileUuid string `json:"fileUuid" binding:"required"`
	}
	var params req

	// 绑定JSON请求体到结构体
	if err := ctx.ShouldBindJSON(&params); err != nil {
		// 绑定失败（如缺少fileUuid、格式错误），返回http.StatusBadRequest错误并提示原因
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "参数错误: " + err.Error()})
		return
	}

	err := r.ReceiveService.UpDateFileStatus(params.FileUuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "文件状态更新失败"})
		return
	}

	ctx.JSON(200, gin.H{"msg": "文件状态已更新"})
}

// QueryReceiveRecords 查询今天接收的记录
func (r *ReceiveController) QueryReceiveRecords(ctx *gin.Context) {
	// 调用服务层方法查询今天接收的记录
	records, err := r.ReceiveService.QueryReceiveRecords()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "查询接收记录失败"})
		return
	}
	// 返回查询结果
	ctx.JSON(http.StatusOK, gin.H{"records": records})
}
