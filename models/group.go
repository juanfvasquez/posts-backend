package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	ImageUrl    string `json:"imageUrl,omitempty"`
	UserId      uint   `json:"user_id"`
}
