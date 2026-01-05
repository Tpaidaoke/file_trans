package repository

import (
	"daoke.com/file_trans/database"
	"daoke.com/file_trans/model"
	"gorm.io/gorm"
	"time"
)

type TransInfoDAO struct {
	db *gorm.DB
}

// NewTransInfoDAO 创建一个新的 TransInfoDAO 实例
func NewTransInfoDAO() *TransInfoDAO {
	return &TransInfoDAO{
		db: database.GetDB(),
	}
}

// Create 保存文件传输信息到数据库
func (t *TransInfoDAO) Create(transInfo *model.TransInfo) error {
	return t.db.Create(transInfo).Error
}

// GetByUUID 根据UUID查询文件信息
func (t *TransInfoDAO) GetByUUID(uuid string) (*model.TransInfo, error) {
	var transInfo model.TransInfo
	result := t.db.Where("file_uuid = ?", uuid).First(&transInfo)
	return &transInfo, result.Error
}

// UpdateSendStatus 标记文件为"已发送"
func (t *TransInfoDAO) UpdateSendStatus(fileUUID string) error {
	return t.db.Model(&model.TransInfo{}).
		Where("file_uuid = ?", fileUUID).
		Update("send_status", true). // 发送成功后更新为true
		Error
}

// UpdateReceiveInfo 取件成功后更新状态和时间
func (t *TransInfoDAO) UpdateReceiveInfo(fileUUID string) error {
	return t.db.Model(&model.TransInfo{}).
		Where("file_uuid = ?", fileUUID).
		Updates(map[string]interface{}{
			"receive_status": true,       // 标记为已取件
			"receive_at":     time.Now(), // 记录当前取件时间
		}).Error
}

// QuerySendRecords 查询今天发送的记录
func (t *TransInfoDAO) QuerySendRecords() ([]model.TransInfo, error) {
	var transInfos []model.TransInfo
	// 查询 send_status = 1 并且 created_at 大于等于 今天0点 的记录
	result := t.db.Where("created_at >= ? and send_status = ?", time.Now().Truncate(24*time.Hour), true).Find(&transInfos)
	return transInfos, result.Error
}

// QueryReceiveRecords 查询今天取件的记录
func (t *TransInfoDAO) QueryReceiveRecords() ([]model.TransInfo, error) {
	// 查询 receive_status = 1 并且 created_at 大于等于 今天0点 的记录
	var transInfos []model.TransInfo
	result := t.db.Where("created_at >= ? and receive_status = ?", time.Now().Truncate(24*time.Hour), true).Find(&transInfos)
	return transInfos, result.Error
}
