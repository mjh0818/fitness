package entity

import "time"

type CoachBook struct {
	UserId        string    `gorm:"type:varchar(200)" json:"user_id"`
	CoachId       string    `gorm:"type:varchar(200)" json:"coach_id"`
	CoachBookDesc string    `gorm:"type:varchar(200)" json:"coach_book_desc"`
	StartTime     time.Time `gorm:"type:datetime" json:"start_time"`
	EndTime       time.Time `gorm:"type:datetime" json:"end_time"`
	CoachBooker   string    `gorm:"type:varchar(200)" json:"coach_booker"`
}
