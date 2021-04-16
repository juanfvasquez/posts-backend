package handlers

import (
	"net/http"
	"strconv"
	"errors"

	"github.com/gorilla/mux"

	"../db"
	"../models"
	"../io/response"
	"../io/request"
	ws "../sockets"
)

func GetCommentsByPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var postId int
	var err error
	if postParam, ok := vars["post"]; ok {
		postId, err = strconv.Atoi(postParam)
		if err != nil {
			response.Error("Given ID is not allowed", http.StatusBadRequest, w)
			return
		}
	} else {
		response.Error("", http.StatusNotFound, w)
		return
	}
	limitParam := r.URL.Query()["limit"]
	offsetParam := r.URL.Query()["offset"]
	if len(limitParam[0]) > 0 || len(offsetParam[0]) > 0 {
		var limit int
		var offset int
		var err error
		if len(limitParam) > 0 {
			limit, err = strconv.Atoi(limitParam[0])
			if err != nil {
				response.Error("Wrong value for Limit param", http.StatusBadRequest, w)
				return
			}
		}
		if len(offsetParam) > 0 {
			offset, err = strconv.Atoi(offsetParam[0])
			if err != nil {
				response.Error("Wrong value for Offset param", http.StatusBadRequest, w)
				return
			}
		}
		comments := db.GetCommentsByPostPaginated(uint(postId), limit, offset)
		status := http.StatusOK
		if len(comments) <= 0 {
			status = http.StatusNoContent
		}
		response.Json(comments, status, w)
	} else {
		comments := db.GetCommentsByPost(uint(postId))
		status := http.StatusOK
		if len(comments) <= 0 {
			status = http.StatusNoContent
		}
		response.Json(comments, status, w)
	}
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var body models.Comment
	err := request.Json(r, &body)
	if err != nil {
		response.Error("Malformed body", http.StatusBadRequest, w)
		return
	}
	err = validateComment(body)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	comment := db.CreateComment(body)
	response.Json(comment, http.StatusCreated, w)
	msg := models.Message{Event: "new::comment", PostId: comment.PostId, Comment: comment}
	ws.Broadcast(msg)
}

func GetLastCommentByPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if postParam, ok := vars["post"]; ok {
		postId, err := strconv.Atoi(postParam)
		if err != nil {
			response.Error("Given ID is not allowed", http.StatusBadRequest, w)
			return
		}
		comment := db.GetLastPostComment(uint(postId))
		status := http.StatusOK
		if comment.ID <= 0 {
			status = http.StatusNoContent
		}
		response.Json(comment, status, w)
	} else {
		response.Error("", http.StatusNotFound, w)
		return
	}
}

func validateComment(body models.Comment) (err error){
	if len(body.Comment) <= 0 {
		err = errors.New("Should contain a comment")
		return
	}
	if body.UserId == 0 {
		err = errors.New("Should contain the creator ID")
		return
	}
	if body.PostId == 0 {
		err = errors.New("Should contain the related Post ID")
		return
	}
	return
}
