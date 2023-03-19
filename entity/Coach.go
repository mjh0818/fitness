package entity

import "gorm.io/gorm"

// 教练表
type Coach struct {
	gorm.Model
	CoachName string `gorm:"type:varchar(200); not null" json:"coach_name"`
	PassWord  string `gorm:"type:varchar(200); not null"  json:"password"`
	Avatar    string `gorm:"type:varchar(200)" json:"avatar"` //头像
	Age       int64  `gorm:"type:int; default:20" json:"age"` //年龄
	Gender    string `gorm:"type:varchar(200)" json:"gender"` //性别
	Mobile    string `gorm:"type:varchar(200); default:11; not null" json:"mobile"`
	Type      string `gorm:"type:varchar(200)" json:"type"`              //执教类型
	State     int64  `gorm:"type:int; default:0; not null" json:"state"` // 0==空闲 1==培训中
	CoachFee  string `gorm:"type:varchar(200)" json:"coach_fee"`         //教练费用
}
