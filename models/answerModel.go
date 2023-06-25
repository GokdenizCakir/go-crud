package models

import "gorm.io/gorm"

type Answer struct{
	gorm.Model
	Value string `json:"value"`
	ApplicantID uint `json:"applicantID"`
	QuestionID uint `json:"questionID"`
}