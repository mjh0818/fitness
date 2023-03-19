package service

import (
	"errors"
	"fitness/dao"
	"fitness/entity"
	"gorm.io/gorm"
)

// 判断用户是否存在
func IsMobileExists(mobile string) bool {
	var user entity.User
	dao.DB.Where("mobile = ?", mobile).First(&user)
	return user.ID != 0
}

// 新建用户
func CreateUser(user *entity.User) (err error) {
	if err = dao.DB.Create(&user).Error; err != nil {
		return errors.New("创建用户失败")
	}
	return nil
}

// 查询用户列表
func GetAllUser(pageNum int, pageSize int) (int, []entity.User) {
	var users []entity.User
	// 计算偏移量
	offset := (pageNum - 1) * pageSize
	// 查询所有的user
	result := dao.DB.Offset(offset).Limit(pageSize).Find(&users)
	// 查不到数据时
	if result.RowsAffected == 0 {
		return 0, nil
	}
	// 获取user总数
	total := len(users)
	// 查询数据
	result.Offset(offset).Limit(pageSize).Find(&users)
	return total, users
}

// 删除用户 根据id
func DeleteUserById(id string) (err error) {
	if err = dao.DB.Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
		return errors.New("删除用户失败")
	}
	return nil
}

// 查询单个用户 根据id
func GetUserById(id string) (user *entity.User, err error) {
	if err = dao.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, errors.New("查询此用户失败")
	}
	return
}

// 更新用户信息
func UpdateUser(id string, user *entity.User) (tx *gorm.DB) {
	tx = dao.DB.Where("id = ?", id).Updates(&user)
	return tx
}

// 用户登录
func UserLogin(mobile string, password string) (user entity.User) {
	dao.DB.Where("mobile = ?", mobile).First(&user)
	return user
}

// 查询密码
func UserPassWord(mobile string) (password string) {
	var user entity.User
	dao.DB.Where("mobile = ?", mobile).First(&user)
	return user.PassWord
}
