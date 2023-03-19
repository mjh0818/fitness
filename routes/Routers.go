package routes

import (
	"fitness/controller"
	"fitness/middleware"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	UserGroup := r.Group("user")
	{
		UserGroup.POST("/login", controller.UserLogin)
		UserGroup.POST("/create", controller.CreateUser)
	}

	UserGroup = r.Group("user")
	UserGroup.Use(middleware.JWT())
	{
		//UserGroup.GET("/getallusers", controller.GetAllUser)
		UserGroup.GET("/getuserbyid", controller.GetUserById)
		UserGroup.GET("/deluserbyid", controller.DeleteUserById)
		UserGroup.PUT("/update", controller.UpdateUser)
	}

	CoachGroup := r.Group("coach")
	{
		CoachGroup.POST("/create", controller.CreateCoach)
		CoachGroup.GET("/delcoachbyid", controller.DeleteCoachById)
		CoachGroup.GET("/delcoachbymobile", controller.DeleteCoachByMobile)
		CoachGroup.PUT("/update", controller.UpdateCoach)
		CoachGroup.GET("/getcoachbyid", controller.GetCoachById)
		CoachGroup.GET("/getcoachbymobile", controller.GetCoachByMobile)
		//CoachGroup.GET("/getcoachs", controller.GetCoachs)

	}

	AreaGroup := r.Group("area")
	{
		AreaGroup.POST("/create", controller.CreateArea)
		AreaGroup.GET("/delareabyid", controller.DeleteAreaById)
		AreaGroup.PUT("/update", controller.UpdateArea)
	}

	return r
}
