package entity

import "gorm.io/gorm"

//场地表

type Area struct {
	gorm.Model
	AreaId    string `gorm:"type:varchar(200)" json:"area_id"`           //场地id
	AreaName  string `gorm:"type:varchar(200)" json:"area_name"`         // 场地名称
	AreaLocal string `gorm:"type:varchar(200)" json:"area_local"`        //场地位置
	AreaDesc  string `gorm:"type:varchar(200)" json:"area_desc"`         // 场地描述
	State     int64  `gorm:"type:int; default:0; not null" json:"state"` //场地状态 0== 空闲   1== 使用中
	AreaFee   string `gorm:"type:varchar(200)" json:"area_fee"`          //场地费
}
