package controller

import (
	"fitness/entity"
	"fitness/service"
	"fitness/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//注册教练

func CreateCoach(c *gin.Context) {
	var coach entity.Coach
	if err := c.BindJSON(&coach); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	if len(coach.Mobile) != 11 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Mobile phone number must be 11 digits",
		})
		return
	}
	if len(coach.PassWord) > 12 || len(coach.PassWord) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "The password is between 6 and 12 digits",
		})
		return
	}
	if coach.CoachName == "" {
		coach.CoachName = "新教练"
	}
	if coach.Gender == "" {
		coach.Gender = "男"
	}

	//判断用户是否存在
	if exists := service.CoachIsMobileExists(coach.Mobile); exists == true {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "该教练已存在",
		})
		return
	}

	//密码加密
	hashpwd := utils.GetHashPwd(coach.PassWord)
	coach.PassWord = string(hashpwd)
	if err := service.CreateCoach(&coach); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "service error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":       "create user success",
			"code":      200,
			"coach_id:": coach.CoachName,
		})
		return
	}
}

//删除教练

func DeleteCoachById(c *gin.Context) {
	var coach entity.Coach
	if err := c.ShouldBindJSON(&coach); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	if err := service.DeleteCoachById(strconv.FormatUint(uint64(coach.ID), 10)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "delete coach error",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "delete coach success",
			"code": 200,
		})
		return
	}
}
func DeleteCoachByMobile(c *gin.Context) {
	var coach entity.Coach
	if err := c.BindJSON(&coach); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	if err := service.DeleteCoachByMobile(coach.Mobile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "delete coach error",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "delete coach success",
			"code": 200,
		})
		return
	}
}

//修改个人信息

func UpdateCoach(c *gin.Context) {
	var coach entity.Coach
	if err := c.BindJSON(&coach); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	tx := service.UpdateCoach(strconv.FormatUint(uint64(coach.ID), 10), &coach)
	if tx.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "update coach success",
			"code": 200,
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "update coach error",
			"code": 400,
		})
	}
}

//查询教练

func GetCoachById(c *gin.Context) {
	var coach entity.Coach
	if err := c.BindJSON(&coach); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	todoList, err := service.GetCoachById(strconv.FormatUint(uint64(coach.ID), 10))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "service error",
		})
		return
	} else if todoList.Mobile == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "The coach does not exist",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "get coach success",
		"code": 200,
		"data": todoList,
	})
	return
}
func GetCoachByMobile(c *gin.Context) {
	var coach entity.Coach
	if err := c.ShouldBindJSON(&coach); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	if todoList, err := service.GetCoachByMobile(coach.Mobile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "service error",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "get coach success",
			"code": 200,
			"data": todoList,
		})
		return
	}
}
func GetCoachs(c *gin.Context) {
	var coach *entity.Coach
	if err := c.ShouldBind(&coach); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "should bind error",
		})
		return
	}
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	pagenum, _ := strconv.Atoi(c.Query("pagenum"))
	list, tx := service.GetAllCoach(pagesize, pagenum)
	if tx.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "get all coach success",
			"code": 200,
			"data": list,
		})
		return
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "not found coach list",
		})
		return
	}

}

//登录

func CoachLogin(c *gin.Context) {
	var coach entity.Coach
	if err := c.BindJSON(&coach); err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"msg": "bind json error",
		})
		return
	}
	if exists := service.IsMobileExists(coach.Mobile); exists != true {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "该用户不存在",
		})
		return
	}

	user := service.UserLogin(coach.Mobile, coach.PassWord)
	if utils.ComparePwd(user.PassWord, coach.PassWord) {
		token, _ := utils.GenerateToken(coach.Mobile, coach.PassWord)
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
