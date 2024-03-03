package routes

import (
	"fmt"
	"strconv"

	"github.com/TeeSSK/GoSimpleREST/model"
	services "github.com/TeeSSK/GoSimpleREST/services"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userService services.UserService
}

func NewUserRoutes(userService services.UserService) *UserRoutes {
	return &UserRoutes{userService}
}

func (r *UserRoutes) SetupRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/", r.GetUsers)
		userGroup.GET("/:id", r.GetUserById)
		userGroup.POST("/", r.CreateUser)
		userGroup.PUT("/:id", r.UpdateUser)
		userGroup.DELETE("/:id", r.DeleteUser)
	}
}

func (r *UserRoutes) GetUsers(c *gin.Context) {
	users, err := r.userService.GetUsers()
	fmt.Println("GetUsers")
	fmt.Println(users)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

func (r *UserRoutes) GetUserById(c *gin.Context) {
	id := c.Param("id")
	uintId, _ := strconv.ParseUint(id, 10, 32)
	user, err := r.userService.GetUserById(uint(uintId))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func (r *UserRoutes) CreateUser(c *gin.Context) {
	var request model.CreateUserRequest
	fmt.Println("CreateUser")
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := r.userService.CreateUser(request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func (r *UserRoutes) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	uintId, _ := strconv.ParseUint(id, 10, 32)
	var request model.UpdateUserRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := r.userService.UpdateUser(uint(uintId), request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func (r *UserRoutes) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	uintId, _ := strconv.ParseUint(id, 10, 32)
	user, err := r.userService.DeleteUser(uint(uintId))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}
