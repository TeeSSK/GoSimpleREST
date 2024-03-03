package main

import (
	"log"

	"github.com/TeeSSK/GoSimpleREST/config"
	"github.com/TeeSSK/GoSimpleREST/repo"
	"github.com/TeeSSK/GoSimpleREST/routes"
	"github.com/TeeSSK/GoSimpleREST/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := config.Connect()
	// routes.UserRoute(router)

	userRepository := repo.NewUserRepo(db)
	userService := services.NewUserService(userRepository)
	userRoute := routes.NewUserRoutes(userService)
	userRoute.SetupRoutes(router)
	router.Run(":8080")
}
