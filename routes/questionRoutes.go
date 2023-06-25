package routes

import (
	"github.com/GokdenizCakir/go-crud/controllers"
	"github.com/GokdenizCakir/go-crud/middlewares"
	"github.com/gin-gonic/gin"
)

func QuestionRoutes(router *gin.Engine) {
	questions := router.Group("/questions") 
	
	questions.Use(middlewares.RequireAuth("admin", "mod"))
	{
		questions.GET("/", controllers.GetAllQuestions)
		questions.POST("/", controllers.CreateQuestion)
	}
	
	questions.Use(middlewares.RequireAuth("admin"))
	{
		questions.DELETE("/:id", controllers.DeleteQuestion)
	}
}