package routes

import (
	"github.com/JoePereiraa/first-crud-golang/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(route *gin.RouterGroup) {
	route.GET("/getUserById/:userId", controller.FindUserById)
	route.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	route.POST("/createUser", controller.CreateUser)
	route.PUT("/updateUser/:userId", controller.UpdateUser)
	route.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
