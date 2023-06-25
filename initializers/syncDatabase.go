package initializers

import "github.com/GokdenizCakir/go-crud/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Answer{})
	DB.AutoMigrate(&models.Question{})
	DB.AutoMigrate(&models.Applicant{})
}