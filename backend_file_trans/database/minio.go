package database

import (
	"context"
	"daoke.com/file_trans/conf"
	"daoke.com/file_trans/logger"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"time"
)

var MinIOClient *minio.Client

// InitMinIO 初始化MinIO客户端
func InitMinIO() {
	minioConf := conf.AppConfig.MinIO

	client, err := minio.New(minioConf.Endpoint, &minio.Options{
		Creds:      credentials.NewStaticV4(minioConf.AccessKey, minioConf.SecretKey, ""),
		Secure:     minioConf.UseSSL,
		MaxRetries: minioConf.MaxRetries,
	})
	if err != nil {
		logger.Error("初始化MinIO客户端失败: %v", err)
	}

	// 检查存储桶是否存在，不存在则创建
	exists, err := client.BucketExists(context.Background(), minioConf.BucketName)
	if err != nil {
		logger.Error("检查存储桶失败: %v", err)
	}
	// 使用带超时的上下文检查存储桶
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	exists, err = client.BucketExists(ctx, minioConf.BucketName)
	if err != nil {
		logger.Error("检查存储桶存在性失败", "bucket", minioConf.BucketName, "error", err)
	}

	if !exists {
		logger.Info("存储桶不存在，开始创建", "bucket", minioConf.BucketName)
		err = client.MakeBucket(ctx, minioConf.BucketName, minio.MakeBucketOptions{})
		if err != nil {
			logger.Error("创建存储桶失败", "bucket", minioConf.BucketName, "error", err)
		}
		logger.Info("存储桶创建成功", "bucket", minioConf.BucketName)
	} else {
		logger.Info("存储桶已存在", "bucket", minioConf.BucketName)
	}

	MinIOClient = client
	logger.Info("MinIO客户端初始化成功")
}
