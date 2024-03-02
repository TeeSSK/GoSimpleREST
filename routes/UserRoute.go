package routes

import (
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
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", r.GetUsers)
		userGroup.GET("/:id", r.GetUserById)
		userGroup.POST("/:id", r.CreateUser)
		userGroup.PUT("/:id", r.UpdateUser)
		userGroup.DELETE("/:id", r.DeleteUser)
	}
}

func (r *UserRoutes) GetUsers(c *gin.Context) {
	users, err := r.userService.GetUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

func (r *UserRoutes) GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := r.userService.GetUserById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func (r *UserRoutes) CreateUser(c *gin.Context) {
	var request model.CreateUserRequest
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
	var request model.UpdateUserRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := r.userService.UpdateUser(id, request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func (r *UserRoutes) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	user, err := r.userService.DeleteUser(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}
