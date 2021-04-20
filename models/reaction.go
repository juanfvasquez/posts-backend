package models

import (
	"gorm.io/gorm"
)

type Reaction struct {
	gorm.Model
	UserId uint `json:"user_id"`
	PostId uint `json:"post_id"`
}
