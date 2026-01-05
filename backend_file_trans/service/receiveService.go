package service

import (
	"context"
	"daoke.com/file_trans/database"
	"daoke.com/file_trans/model"
	"daoke.com/file_trans/repository"
	"daoke.com/file_trans/utils"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type ReceiveService struct {
	transInfoDB *repository.TransInfoDAO
	rClient     *redis.Client // 补充Redis依赖（之前缺失）
}

func NewReceiveService() *ReceiveService {
	return &ReceiveService{
		transInfoDB: repository.NewTransInfoDAO(),
		rClient:     database.RClient, // 从全局获取Redis客户端
	}
}

// GetTransInfoByPickupCode 根据取件码查询并更新取件信息
func (r *ReceiveService) GetTransInfoByPickupCode(pickupCode string) (*model.TransInfo, error) {
	// 1. 从Redis获取fileUuid（取件码映射关系）
	ctx := context.Background()
	fileUuid, err := r.rClient.Get(ctx, pickupCode).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, errors.New("pickupCode not found")
		}
		return nil, err
	}

	// 2. 查询文件信息
	transInfo, err := r.transInfoDB.GetByUUID(fileUuid)
	if err != nil {
		return nil, err
	}

	// 3. 检查是否过期
	if transInfo.IsExpire {
		return nil, errors.New("取件码已过期！")
	}

	// 4. 生成预签名下载链接
	bucketName, objectName := utils.ParseStorageURL(transInfo.StorageUrl)

	// 生成15分钟有效的下载链接（可根据需求调整过期时间）
	preSignedURL, err := utils.GeneratePreSignedDownloadURL(bucketName, objectName, 15*time.Minute)
	if err != nil {
		return nil, errors.New("生成文件下载链接失败!")
	}
	transInfo.StorageUrl = preSignedURL

	return transInfo, nil
}

// UpDateFileStatus 当前端返回文件下载完成，更新文件状态
func (r *ReceiveService) UpDateFileStatus(fileUuid string) error {
	//取件成功，更新取件状态和时间
	if err := r.transInfoDB.UpdateReceiveInfo(fileUuid); err != nil {
		return err
	}
	return nil
}

// QueryReceiveRecords 查询今天接收的记录
func (r *ReceiveService) QueryReceiveRecords() (interface{}, interface{}) {
	// 1. 从数据库查询今天接收的记录
	records, err := r.transInfoDB.QueryReceiveRecords()
	if err != nil {
		return nil, err
	}
	return records, nil
}
