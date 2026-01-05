package model

import (
	"time"
)

type TransInfo struct {
	ID            uint      `gorm:"primaryKey"`
	FileUuid      string    `gorm:"type:varchar(64);not null;unique"`                  // 文件唯一标识
	FileName      string    `gorm:"type:varchar(255);not null"`                        // 文件名
	FileType      string    `gorm:"type:varchar(255);not null"`                        // 文件类型
	FileSize      int64     `gorm:"not null"`                                          // 文件大小（字节）
	StorageUrl    string    `gorm:"type:varchar(512);not null"`                        // MinIO存储地址
	IsExpire      bool      `gorm:"not null;default:0"`                                // 是否过期（0-未过期，1-已过期）
	SendStatus    bool      `gorm:"not null;default:0"`                                // 发送状态（0-未发送，1-已发送）
	ReceiveStatus bool      `gorm:"not null;default:0"`                                // 取件状态（0-未取件，1-已取件）
	ReceiveAt     time.Time `gorm:"type:timestamp;default:NULL"`                       // 取件时间（默认NULL，取件时更新）
	CreatedAt     time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`                // 创建时间（自动记录）
	UpdateAt      time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;autoUpdateTime"` // 自动更新时间（GORM自动维护）
}

func (transInfo *TransInfo) TableName() string {
	return "trans_info"
}
