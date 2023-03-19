package controller

import (
	"fitness/entity"
	"fitness/service"
	"fitness/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateArea(c *gin.Context) {
	var area entity.Area
	if err := c.ShouldBind(&area); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "should bind error",
		})
		return
	}
	area.AreaId = utils.RandInt(12)
	//if area.AreaName == "" {
	//	area.AreaName = "新场地--暂未开放"
	//}
	//if area.AreaDesc == "" {
	//	area.AreaDesc = "无"
	//}
	//if area.AreaLocal == "" {
	//	area.AreaLocal = "无"
	//}
	if err := service.CreateArea(&area); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "service error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":     "create area success",
		"state":   "空闲中",
		"code":    "200",
		"local":   area.AreaLocal,
		"area_id": area.AreaId,
	})
	return
}

func DeleteAreaById(c *gin.Context) {
	var area entity.Area
	if err := c.BindJSON(&area); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "should bind error",
		})
		return
	}
	if err := service.DeleteAreaById(strconv.FormatUint(uint64(area.ID), 10)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "del area service error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "del area success",
		"id":  area.ID,
	})
	return
}

func UpdateArea(c *gin.Context) {
	var area entity.Area
	if err := c.ShouldBind(&area); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "should bind error",
		})
		return
	}
	if err := service.UpdateArea(&area); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "update area service error",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "update area service success",
			"code": 200,
		})
		return
	}
}
