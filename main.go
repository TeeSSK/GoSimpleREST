package main

import (
	"github.com/TeeSSK/GoSimpleREST/config"
	"github.com/TeeSSK/GoSimpleREST/repo"
	"github.com/TeeSSK/GoSimpleREST/routes"
	"github.com/TeeSSK/GoSimpleREST/services"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	db := config.Connect()
	// routes.UserRoute(router)

	userRepository := repo.NewUserRepo(db)
	userService := services.NewUserService(userRepository)
	userRoute := routes.NewUserRoutes(userService)
	userRoute.SetupRoutes(router)
	router.Run(":8080")
}
