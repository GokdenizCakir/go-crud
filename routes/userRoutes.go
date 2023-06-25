package routes

import (
	"github.com/GokdenizCakir/go-crud/controllers"
	"github.com/GokdenizCakir/go-crud/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	users := router.Group("/users") 
	{
		router.Use(middlewares.RequireAuth("admin"))
		users.GET("/", controllers.GetUsers)
	}
}