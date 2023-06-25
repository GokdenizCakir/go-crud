package controllers

import (
	"net/http"

	"github.com/GokdenizCakir/go-crud/initializers"
	"github.com/GokdenizCakir/go-crud/models"
	"github.com/gin-gonic/gin"
)

func GetAllQuestions(c *gin.Context) {
	var questions []models.Question

	err := initializers.DB.Find(&questions).Error;
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get questions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"questionsFound": len(questions), "questions": questions})
}

func CreateQuestion(c *gin.Context) {
	var question models.Question

	err :=c.ShouldBindJSON(&question)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid question format"})
		return
	}

	result := initializers.DB.Create(&question);
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create question"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Question created successfully", "question": question})
}

func DeleteQuestion(c *gin.Context) {
	id := c.Param("id");

	var question models.Question

	result := initializers.DB.Delete(&question, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete question"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Question deleted successfully"})
}