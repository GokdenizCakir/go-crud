package models

import "gorm.io/gorm"

type Applicant struct{
	gorm.Model
	ClientIP string `json:"clientIP" gorm:"unique"`
	Answers []Answer `json:"answers" gorm:"foreignkey:ApplicantID"`
}