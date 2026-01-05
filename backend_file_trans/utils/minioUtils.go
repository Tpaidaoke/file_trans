package utils

import (
	"context"
	"daoke.com/file_trans/database"
	"net/url"
	"strings"
	"time"
)

/*
GeneratePreSignedDownloadURL 生成MinIO文件的预签名下载链接
bucketName: MinIO存储桶名称
objectName: 存储在MinIO中的对象路径（如"files/12345_logo.png"）
expiry: 链接过期时间（如15分钟）
*/
func GeneratePreSignedDownloadURL(bucketName, objectName string, expiry time.Duration) (string, error) {
	ctx := context.Background()
	// 生成预签名URL（包含临时访问凭证）
	preSignedURL, err := database.MinIOClient.PresignedGetObject(
		ctx,
		bucketName,
		objectName,
		expiry,
		url.Values{}, // 额外请求参数（可选）
	)
	if err != nil {
		return "", err
	}
	return preSignedURL.String(), nil
}

// 从StorageUrl中解析bucket和object名称
func ParseStorageURL(storageUrl string) (bucketName, objectName string) {
	parsedURL, err := url.Parse(storageUrl)
	if err != nil {
		return "", ""
	}
	// 路径格式: /{bucketName}/{objectName}，需去除开头的"/"
	pathParts := strings.Split(strings.TrimPrefix(parsedURL.Path, "/"), "/")
	if len(pathParts) < 2 {
		return "", ""
	}
	bucketName = pathParts[0]
	objectName = strings.Join(pathParts[1:], "/") // 拼接剩余部分为object路径
	return bucketName, objectName
}
