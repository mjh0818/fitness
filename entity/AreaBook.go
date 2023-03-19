package entity

import "time"

// 场地预约表
type AreaBook struct {
	UserId       string    `gorm:"type:varchar(200)" json:"user_id"`
	AreaId       string    `gorm:"type:varchar(200)" json:"area_id"`
	AreaBookDesc string    `gorm:"type:varchar(200)" json:"area_book_desc"`
	StartTime    time.Time `gorm:"type:datetime" json:"start_time"`
	EndTime      time.Time `gorm:"type:datetime" json:"end_time"`
	AreaBooker   string    `gorm:"type:varchar(200)" json:"area_booker"`
}
