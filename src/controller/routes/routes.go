package routes

import (
	"github.com/DiogoLuizBonaparte/go_api.git/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(rt *gin.RouterGroup) {
	rt.GET("/getUserById/:userId", controller.FindUserById)
	rt.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	rt.PUT("/updateUser/:userId", controller.UpdateUser)
	rt.POST("/createUser/", controller.CreateUser)
	rt.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
