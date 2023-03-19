package controller

import (
	"fitness/entity"
	"fitness/service"
	"fitness/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {

	var user entity.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	if len(user.Mobile) != 11 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Mobile phone number must be 11 digits",
		})
		return
	}
	if len(user.PassWord) > 12 || len(user.PassWord) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "The password is between 6 and 12 digits",
		})
		return
	}
	if user.NickName == "" {
		user.NickName = "新用户"
	}
	if user.Gender == "" {
		user.Gender = "男"
	}
	user.UserID = utils.RandInt(12)
	//判断用户是否存在
	if exists := service.IsMobileExists(user.Mobile); exists == true {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "该用户已存在",
		})
		return
	}

	//密码加密
	hashpwd := utils.GetHashPwd(user.PassWord)
	user.PassWord = string(hashpwd)
	if err := service.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "service error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":      "create user success",
			"code":     200,
			"user_id:": user.UserID,
		})
		return
	}
}

func GetUserById(c *gin.Context) {
	var user entity.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	if todoList, err := service.GetUserById(strconv.FormatUint(uint64(user.ID), 10)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "service error",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "get user success",
			"code": 200,
			"data": todoList,
		})
		return
	}
}

func GetAllUser(c *gin.Context) {
	var users []entity.User
	if err := c.BindJSON(&users); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	pagenum, _ := strconv.Atoi(c.Query("pagenum"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	total, user := service.GetAllUser(pagenum, pagesize)
	if total == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "未查询到数据",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "查询成功",
			"data": user,
		})
	}
}

func DeleteUserById(c *gin.Context) {
	var users entity.User
	if err := c.BindJSON(&users); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}

	if err := service.DeleteUserById(strconv.FormatUint(uint64(users.ID), 10)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "delete user error",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "delete user success",
			"code": 200,
		})
		return
	}
}

func UpdateUser(c *gin.Context) {
	var users entity.User
	if err := c.BindJSON(&users); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	tx := service.UpdateUser(strconv.FormatUint(uint64(users.ID), 10), &users)
	if tx.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "update user success",
			"code": 200,
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "update user error",
			"code": 400,
		})
	}
}

func UserLogin(c *gin.Context) {
	var users entity.User
	if err := c.BindJSON(&users); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	if exists := service.IsMobileExists(users.Mobile); exists != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "该用户不存在",
		})
		return
	}

	user := service.UserLogin(users.Mobile, users.PassWord)
	if utils.ComparePwd(user.PassWord, users.PassWord) {
		token, _ := utils.GenerateToken(users.Mobile, users.PassWord)
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "登录成功",
			"token":   token,
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "密码错误",
		})
		return
	}
}
