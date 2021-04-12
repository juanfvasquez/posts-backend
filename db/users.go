package db

import (
	"strings"
	"math/rand"
	"time"

	"../models"
)

var urls = []string{
	"https://raw.githubusercontent.com/Ashwinvalento/cartoon-avatar/master/lib/images/male/45.png",
	"https://raw.githubusercontent.com/Ashwinvalento/cartoon-avatar/master/lib/images/male/86.png",
	"https://raw.githubusercontent.com/Ashwinvalento/cartoon-avatar/master/lib/images/male/5.png",
}

func GetUsers() (users []models.User) {
	db := GetConnection()
	db.Find(&users)
	return
}

func GetUser(id uint) (user models.User) {
	db := GetConnection()
	db.Find(&user, id)
	return
}

func SearchUser(searchParam string) (users []models.User) {
	db := GetConnection()
	str := strings.ToLower(searchParam)
	db.Find(&users, "LOWER(name) LIKE ? OR LOWER(email) LIKE ?", str, str)
	return
}

func CreateUser(body models.User) (user models.User) {
	db := GetConnection()
	body.Avatar = getRandomAvatar()
	db.Create(&body)
	db.Last(&user)
	return
}

func getRandomAvatar() (avatar string) {
	rand.Seed(time.Now().UnixNano())
	avatar = urls[rand.Intn(len(urls))]
	return
}

func GetUserByEmail(email string) (user models.User) {
	db := GetConnection()
	db.Find(&user, "email = ?", email)
	return
}