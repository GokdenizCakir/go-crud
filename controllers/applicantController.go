package controllers

import (
	"net/http"

	"github.com/GokdenizCakir/go-crud/initializers"
	"github.com/GokdenizCakir/go-crud/models"

	"github.com/gin-gonic/gin"
)

func CreateApplicant(c *gin.Context) {
	var newApplicant models.Applicant;
	
	var request struct {
		Answers []models.Answer `json:"answers"`
	}
	err := c.ShouldBindJSON(&request)
	
	if err != nil || len(request.Answers) < 1  {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid answers payload"})
        return
	}

	var existingApplicant models.Applicant;
	if err := initializers.DB.Where("clientIP = ?", c.ClientIP()).First(&existingApplicant).Error; err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "User with the same IP already exists"})
        return
    }
	
	newApplicant.ClientIP = c.ClientIP()
    newApplicant.Answers = request.Answers
	
	errCreate := initializers.DB.Create(&newApplicant).Error;
	if errCreate != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errCreate})
        return
    }
	
	c.JSON(http.StatusCreated, newApplicant)
	
}


func GetAllApplicants(c *gin.Context) {
	var applicants []models.Applicant;

	result := initializers.DB.Find(&applicants);
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": applicants})
}

func GetApplicant(c *gin.Context) {
	var applicant models.Applicant;
	id := c.Param("id");

	err := initializers.DB.Where("id = ?", id).Preload("Answers").First(&applicant).Error;
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Applicant not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": applicant})
}