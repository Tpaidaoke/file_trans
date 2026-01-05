package controller

import (
	"daoke.com/file_trans/logger"
	"daoke.com/file_trans/model"
	"daoke.com/file_trans/service"
	"daoke.com/file_trans/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type SendController struct {
	SendService *service.SendService
}

// NewSendController 创建一个新的 SendController 实例
func NewSendController() *SendController {
	return &SendController{
		SendService: service.NewSendService(),
	}
}

// Send 处理发送文件的请求
func (s *SendController) Send(ctx *gin.Context) {
	// 获取请求类型：文本或文件
	transType := ctx.PostForm("type")

	// 获取前端传入的过期时间参数
	expireTipStr := ctx.PostForm("expireTip")
	expireUnit := ctx.PostForm("expireUnit")

	// 验证参数存在性
	if expireTipStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请传入过期时间数值expireTip"})
		return
	}
	if expireUnit == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "请传入过期时间单位expireUnit"})
		return
	}

	// 验证数值有效性
	expireTip, err := strconv.Atoi(expireTipStr)
	if err != nil || expireTip <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "过期时间数值必须为正整数"})
		return
	}

	// 验证单位有效性并计算过期时间
	var expiry time.Duration
	var expiresIn int
	switch expireUnit {
	case "分钟":
		expiry = time.Duration(expireTip) * time.Minute
		expiresIn = expireTip * 60
	case "小时":
		expiry = time.Duration(expireTip) * time.Hour
		expiresIn = expireTip * 3600
	case "天":
		expiry = time.Duration(expireTip) * 24 * time.Hour
		expiresIn = expireTip * 86400
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "过期时间单位无效"})
		return
	}

	var (
		fileURL         string
		fileDownloadURL string
		fileName        string
		fileUUID        string
		fileSize        int64
	)

	// 根据类型处理
	if transType == "text" {
		// 处理文本类型
		textContent := ctx.PostForm("text")
		if textContent == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "文本内容不能为空"})
			return
		}

		// 验证文本内容长度
		maxTextSize := 10 * 1024 * 1024 // 10MB
		if len(textContent) > maxTextSize {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "文本内容过大，最大支持10MB"})
			return
		}

		// 生成文件名
		fileName = fmt.Sprintf("文本_%d.txt", time.Now().UnixNano())
		fileSize = int64(len(textContent))

		// 将文本内容转换为Reader
		reader := strings.NewReader(textContent)

		// 调用统一上传方法，指定类型为"text"
		fileURL, err = s.SendService.UploadToMinIO(fileName, fileSize, reader, "text")
		if err != nil {
			logger.Error("上传文本文件失败", "err", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "处理文本失败"})
			return
		}

	} else if transType == "file" {
		// 处理文件类型
		form, err := ctx.MultipartForm()
		if err != nil {
			logger.Warn("获取上传文件失败", "err", err, "client_ip", ctx.ClientIP())
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "获取文件失败"})
			return
		}

		files := form.File["files"]
		if len(files) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "未找到文件"})
			return
		}

		// 判断文件数量，超过2个则压缩
		if len(files) >= 2 {
			fileMap := make(map[string]io.Reader)
			fileSizeTotal := int64(0)

			// 收集所有文件
			for _, file := range files {
				src, err := file.Open()
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": "打开文件失败"})
					return
				}
				defer src.Close()

				fileMap[file.Filename] = src
				fileSizeTotal += file.Size
			}

			// 压缩文件
			zipReader, zipSize, err := utils.CompressFiles(fileMap)
			if err != nil {
				logger.Error("压缩文件失败", "err", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "压缩文件失败"})
				return
			}

			fileName = fmt.Sprintf("files_%d.zip", time.Now().UnixNano())
			fileSize = zipSize

			// 调用统一上传方法，指定类型为"file"
			fileURL, err = s.SendService.UploadToMinIO(fileName, fileSize, zipReader, "file")
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "文件上传失败"})
				return
			}
		} else {
			// 单个文件直接上传
			file := files[0]
			src, err := file.Open()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "打开文件失败"})
				return
			}
			defer src.Close()

			fileName = file.Filename
			fileSize = file.Size

			// 调用统一上传方法，指定类型为"file"
			fileURL, err = s.SendService.UploadToMinIO(fileName, fileSize, src, "file")
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "文件上传失败"})
				return
			}
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的传输类型"})
		return
	}

	// 为所有类型生成预签名下载链接
	bucketName, objectName := utils.ParseStorageURL(fileURL)
	fileDownloadURL, err = utils.GeneratePreSignedDownloadURL(bucketName, objectName, 15*time.Minute)
	if err != nil {
		logger.Error("生成文件下载链接失败", "err", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "生成文件下载链接失败!"})
		return
	}

	// 生成唯一标识
	fileUUID = s.SendService.GenerateFileUUID()
	accessKey := ctx.PostForm("pickupCode")

	// 保存文件信息到数据库
	transInfo := &model.TransInfo{
		FileUuid:   fileUUID,
		FileName:   fileName,
		FileSize:   fileSize,
		StorageUrl: fileURL,
		FileType:   transType, // 保存类型标识
	}

	if err := s.SendService.SaveToDB(transInfo); err != nil {
		logger.Error("保存文件信息到数据库失败", "err", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "保存文件信息失败"})
		return
	}

	// 保存到Redis
	if err := s.SendService.SaveToRedis(accessKey, fileUUID, expiry); err != nil {
		logger.Error("保存文件信息到Redis失败", "err", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "保存访问密钥失败"})
		return
	}

	// 更新文件发送状态
	if err := s.SendService.UpdateSendStatus(fileUUID); err != nil {
		logger.Error("更新文件发送状态失败", "err", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "更新文件发送状态失败"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":             "发送成功",
		"fileName":        fileName,
		"fileSize":        fileSize,
		"fileDownloadURL": fileDownloadURL,
		"fileUuid":        fileUUID,
		"accessKey":       accessKey,
		"expiresIn":       expiresIn,
		"unit":            expireUnit,
		"type":            transType, // 返回类型，前端可能需要
	})
}

// QuerySendRecords 查询发送记录
func (s *SendController) QuerySendRecords(context *gin.Context) {
	//查询当天的发送记录
	records, err := s.SendService.QuerySendRecords()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "查询发送记录失败"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"records": records})
}
