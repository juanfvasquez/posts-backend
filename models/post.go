package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	PostedText string     `json:"posted_text"`
	ImageUrl   string     `json:"image_url,omitempty"`
	UserId     uint       `json:"user_id"`
	User       User       `gorm:"foreignKey:UserId" json:"user,omitempty"`
	Comments   []Comment  `gorm:"foreignKey:PostId" json:"comments,omitempty"`
	Reactions  []Reaction `gorm:"foreignKey:PostId" json:"reactions,omitempty"`
}
