package routes

import (
	"github.com/GokdenizCakir/go-crud/controllers"
	"github.com/GokdenizCakir/go-crud/middlewares"

	"github.com/gin-gonic/gin"
)

func ApplicantRoutes(router *gin.Engine) {
	applicant := router.Group("/applicant") 
	applicant.POST("/", controllers.CreateApplicant)



	applicant.Use(middlewares.RequireAuth("admin", "mod"))
	{
		applicant.GET("/:id", controllers.GetApplicant)
		applicant.GET("/", controllers.GetAllApplicants)
	}
}