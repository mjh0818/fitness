package entity

import "gorm.io/gorm"

// 用户表
type User struct {
	gorm.Model
	UserID   string `gorm:"type:varchar(200)" json:"user_id"`   //userid
	NickName string `gorm:"type:varchar(200)" json:"nick_name"` //用户名
	PassWord string `gorm:"type:varchar(200)" json:"password" validate:"required,max=12,min=6"`
	Gender   string `gorm:"type:varchar(200)" json:"gender"`
	Mobile   string `gorm:"type:varchar(200)" json:"mobile" validate:"required"`
}
