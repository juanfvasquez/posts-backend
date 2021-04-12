package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Comment string 	`json:"comment"`
	UserId 	uint 		`json:"user_id"`
	User 		User 		`gorm:"foreignKey:UserId"`
	PostId 	uint		`json:"post_id"`
}