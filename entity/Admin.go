package entity

import (
	"gorm.io/gorm"
)

// 管理员表
type Admin struct {
	gorm.Model
	AdminId  string `gorm:"type:varchar(200)" json:"adminid"`
	Mobile   string `gorm:"type:varchar(200); not null" json:"mobile"`
	PassWord string `gorm:"type:varchar(200)" json:"password"`
}
