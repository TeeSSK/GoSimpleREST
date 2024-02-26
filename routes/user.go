package routes

import (
	"github.com/TeeSSK/GoSimpleREST/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/user", controller.GetUsers)
	router.POST("/user", controller.CreateUser)
	router.PUT("/user/:id", controller.UpdateUser)
	router.DELETE("/user/:id", controller.DeleteUser)
}
