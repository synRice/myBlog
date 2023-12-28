package models

import (
	"time"

	"gorm.io/gorm"
)

type TestModel struct {
	ID        int            `gorm:"column:id;primarykey"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
	Username  string         `gorm:"column:username;size:255;Index"`
	Content   string         `gorm:"column:content;size:255"`
}

// 设置模型对应的表名
func (TestModel) TableName() string {
	return "test1"
}
