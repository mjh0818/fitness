package service

import (
	"errors"
	"fitness/dao"
	"fitness/entity"
	"gorm.io/gorm"
)

// 判断教练是否存在

func CoachIsMobileExists(mobile string) bool {
	var coach entity.Coach
	dao.DB.Where("mobile = ?", mobile).First(&coach)
	return coach.ID != 0
}

// 新建教练

func CreateCoach(coach *entity.Coach) (err error) {
	if err = dao.DB.Create(&coach).Error; err != nil {
		return errors.New("创建教练失败")
	}
	return nil
}

// 删除教练

func DeleteCoachById(id string) (err error) {
	if err := dao.DB.Where("id = ?", id).Delete(&entity.Coach{}).Error; err != nil {
		return errors.New("删除教练失败")
	}
	return
}
func DeleteCoachByMobile(mobile string) (err error) {
	if err := dao.DB.Where("mobile = ?", mobile).Delete(&entity.Coach{}).Error; err != nil {
		return errors.New("删除教练失败")
	}
	return
}

// 查询教练

func GetCoachById(id string) (coach *entity.Coach, err error) {
	if err = dao.DB.Where("id = ?", id).First(&coach).Error; err != nil {
		return nil, errors.New("查询此教练失败")
	}
	return
}
func GetCoachByMobile(mobile string) (coach *entity.Coach, err error) {
	if err = dao.DB.Where("mobile = ?", mobile).Find(&coach).Error; err != nil {
		return nil, errors.New("查询此教练失败")
	}
	return
}

// 查询教练列表

func GetAllCoach(PageSize int, PageNum int) (coachList []entity.Coach, tx *gorm.DB) {
	tx = dao.DB.Limit(PageSize).Offset((PageNum - 1) * PageSize).Find(&coachList)
	return
}

// 更新教练信息

func UpdateCoach(id string, coach *entity.Coach) (tx *gorm.DB) {
	tx = dao.DB.Where("id = ?", id).Updates(&coach)
	return tx
}
