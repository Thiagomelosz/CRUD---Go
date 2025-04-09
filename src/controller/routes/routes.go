package routes

import (
	"github.com/Thiagomelosz/CRUD---Go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.createUser)
	r.PUT("/updateUser/:userId")
	r.DELETE("/deleteUser/:userId")
}
