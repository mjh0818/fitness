package service

import (
	"errors"
	"fitness/dao"
	"fitness/entity"
)

// 增加场地
func CreateArea(area *entity.Area) (err error) {
	if err := dao.DB.Create(&area).Error; err != nil {
		return errors.New("创建场地失败")
	}
	return nil
}

// 删除场地根据id
func DeleteAreaById(id string) (err error) {
	var area *entity.Area
	if err := dao.DB.Where("id = ?", id).Delete(&area).Error; err != nil {
		return errors.New("del area dao error")
	}
	return nil
}

// 修改场地信息
func UpdateArea(area *entity.Area) error {
	err := dao.DB.Where("id = ?", area.ID).Updates(&area).Error
	if err != nil {
		return errors.New("更新场地信息失败")
	}
	return nil
}
