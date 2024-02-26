package main

import (
	"github.com/TeeSSK/GoSimpleREST/config"
	"github.com/TeeSSK/GoSimpleREST/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")
}
