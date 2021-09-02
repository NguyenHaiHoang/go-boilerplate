package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID                  uint64             `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt           *time.Time         `gorm:"column:created_at;autoCreateTime" json:"-"`
	CreatedBy           *uint64            `gorm:"column:created_by" json:"-"`
	UpdatedAt           *time.Time         `gorm:"column:updated_at;autoUpdateTime" json:"-"`
	UpdatedBy           *uint64            `gorm:"column:updated_by" json:"-"`
	DeletedAt           *gorm.DeletedAt    `gorm:"column:deleted_at" json:"-"`
}
