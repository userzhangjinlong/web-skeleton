package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	//创建时间
	CreateAt string `gorm:"column:createAt" json:"createAt"`
	//更新时间
	UpdateAt string `gorm:"column:updateAt" json:"updateAt"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) error {
	m.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	m.UpdateAt = time.Now().Format("2006-01-02 15:04:05")

	return nil
}

func (m *Model) BeforeUpdate(tx *gorm.DB) error {
	m.UpdateAt = time.Now().Format("2006-01-02 15:04:05")

	return nil
}

//Paginate 自定义分页
func (m *Model) Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
