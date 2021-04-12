package db

import (
	"gorm.io/gorm"

	"../models"
)

const COMMENTS_PER_POST = 1

func GetPosts() (posts []models.Post) {
	db := GetConnection()
	db.Preload("User").Preload("Comments", preloadLatestComment(COMMENTS_PER_POST)).Order("created_at DESC").Find(&posts)
	return
}

func GetPostsByUser(userId uint) (posts []models.Post) {
	db := GetConnection()
	db.Preload("User").Preload("Comments.User", preloadLatestComment(COMMENTS_PER_POST)).Order("created_at DESC").Find(&posts, "user_id = ?", userId)
	return
}

func GetPostsPaginated(limit, offset int) (posts []models.Post) {
	db := GetConnection()
	db.Preload("User").Preload("Comments.User", preloadLatestComment(COMMENTS_PER_POST)).Order("created_at DESC").Limit(limit).Offset(offset).Find(&posts)
	return
}

func GetPost(id uint) (post models.Post) {
	db := GetConnection()
	db.Preload("User").Preload("Comments.User", preloadLatestComment(COMMENTS_PER_POST)).Order("created_at DESC").Find(&post, id)
	return
}

func CreatePost(body models.Post) (post models.Post) {
	db := GetConnection()
	db.Create(&body)
	db.Last(&post)
	return
}

func preloadLatestComment(num int) func (*gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		return db.Order("created_at DESC").Limit(num)
	}
}