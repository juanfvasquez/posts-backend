package db

import (
	"../models"
)

func GetCommentsByPost(postId uint) (comments []models.Comment) {
	db := GetConnection()
	db.Find(&comments, "post_id = ?", postId)
	return
}

func GetCommentsByPostPaginated(postId uint, limit, offset int) (comments []models.Comment) {
	db := GetConnection()
	db.Limit(limit).Offset(offset).Find(&comments, "post_id = ?", postId)
	return
}

func CreateComment(body models.Comment) (comment models.Comment) {
	db := GetConnection()
	db.Create(&body)
	db.Last(&comment)
	return
}

func GetLastPostComment(postId uint) (comment models.Comment) {
	db := GetConnection()
	db.Last(&comment, "post_id = ?", postId)
	return
}