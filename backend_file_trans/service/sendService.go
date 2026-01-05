package service

import (
	"context"
	"daoke.com/file_trans/conf"
	"daoke.com/file_trans/database"
	"daoke.com/file_trans/logger"
	"daoke.com/file_trans/model"
	"daoke.com/file_trans/repository"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"io"
	"time"
)

type SendService struct {
	transInfoDB *repository.TransInfoDAO
}

// NewSendService 创建一个新的 SendService 实例
func NewSendService() *SendService {
	return &SendService{
		transInfoDB: repository.NewTransInfoDAO(),
	}
}

// UploadToMinIO 通用上传方法，支持文本和文件
func (s *SendService) UploadToMinIO(fileName string, fileSize int64, reader io.Reader, fileType string) (string, error) {
	bucketName := conf.AppConfig.MinIO.BucketName

	// 根据类型决定存储路径和文件名
	var objName string
	var contentType string

	if fileType == "text" {
		// 文本文件：存储到texts目录，固定名称
		objName = fmt.Sprintf("texts/%s", fileName)
		contentType = "text/plain; charset=utf-8"
	} else {
		// 普通文件：存储到files目录，添加时间戳避免重名
		objName = fmt.Sprintf("files/%d_%s", time.Now().UnixNano(), fileName)
		contentType = "application/octet-stream"
	}

	logger.Debug("开始上传文件到MinIO",
		"bucket", bucketName,
		"objName", objName,
		"type", fileType,
		"size", fileSize)

	// 上传文件
	_, err := database.MinIOClient.PutObject(
		context.Background(),
		bucketName,
		objName,
		reader,
		fileSize,
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		logger.Error("MinIO文件上传失败", "err", err, "file", fileName, "type", fileType)
		return "", err
	}

	// 生成文件URL
	endpoint := conf.AppConfig.MinIO.Endpoint
	var fileURL string
	if conf.AppConfig.MinIO.UseSSL {
		fileURL = fmt.Sprintf("https://%s/%s/%s", endpoint, bucketName, objName)
	} else {
		fileURL = fmt.Sprintf("http://%s/%s/%s", endpoint, bucketName, objName)
	}

	logger.Info("MinIO文件上传成功",
		"fileURL", fileURL,
		"fileName", fileName,
		"type", fileType)

	return fileURL, nil
}

// SaveToRedis 保存文件信息到Redis
func (s *SendService) SaveToRedis(accessKey, fileUUID string, expiry time.Duration) error {
	ctx := context.Background()
	return database.RClient.Set(ctx, accessKey, fileUUID, expiry).Err()
}

// SaveToDB 保存文件信息到数据库
func (s *SendService) SaveToDB(transInfo *model.TransInfo) error {
	return s.transInfoDB.Create(transInfo)
}

// GenerateFileUUID 生成文件UUID
func (s *SendService) GenerateFileUUID() string {
	return uuid.New().String()
}

// QuerySendRecords 查询今天发送的记录
func (s *SendService) QuerySendRecords() ([]model.TransInfo, error) {
	return s.transInfoDB.QuerySendRecords()
}

// UpdateSendStatus 更新文件发送状态
func (s *SendService) UpdateSendStatus(fileUUID string) error {
	return s.transInfoDB.UpdateSendStatus(fileUUID)
}
