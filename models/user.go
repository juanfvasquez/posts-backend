package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name 			string 		`json:"name"`
	Email 		string 		`json:"email"`
	Password 	string		`json:"password,omitempty"`
	Avatar 		string 		`json:"avatar"`
	Posts 		[]Post 		`gorm:"foreignKey:UserId" json:"posts,omitempty"`
	Comments 	[]Comment `gorm:"foreignKey:UserId" json:"comments,omitempty"`
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	u.Password = ""
	return
}