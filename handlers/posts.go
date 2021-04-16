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

func GetPosts(w http.ResponseWriter, r *http.Request) {
	limitParam := r.URL.Query()["limit"]
	offsetParam := r.URL.Query()["offset"]
	if len(limitParam) > 0 || len(offsetParam) > 0 {
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
		posts := db.GetPostsPaginated(limit, offset)
		status := http.StatusOK
		if len(posts) <= 0 {
			status = http.StatusNoContent
		}
		response.Json(posts, status, w)
	} else {
		posts := db.GetPosts()
		status := http.StatusOK
		if len(posts) <= 0 {
			status = http.StatusNoContent
		}
		response.Json(posts, status, w)
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var body models.Post
	err := request.Json(r, &body)
	if err != nil {
		response.Error("Malformed body", http.StatusBadRequest, w)
		return
	}
	err = validatePost(body)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	post := db.CreatePost(body)
	response.Json(post, http.StatusCreated, w)
	msg := models.Message{Event: "new::post", Post: post}
	ws.Broadcast(msg)
}

func GetPostsByUser(w http.ResponseWriter, r * http.Request) {
	vars := mux.Vars(r)
	if userParam, ok := vars["user"]; ok {
		userId, err := strconv.Atoi(userParam)
		if err != nil {
			response.Error("Given ID is not allowed", http.StatusBadRequest, w)
			return
		}
		posts := db.GetPostsByUser(uint(userId))
		status := http.StatusOK
		if len(posts) <= 0 {
			status = http.StatusNoContent
		}
		response.Json(posts, status, w)
	} else {
		response.Error("", http.StatusNotFound, w)
	}
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if idParam, ok := vars["id"]; ok {
		id, err := strconv.Atoi(idParam)
		if err != nil {
			response.Error("Given ID is not allowed", http.StatusBadRequest, w)
			return
		}
		post := db.GetPost(uint(id))
		status := http.StatusOK
		if post.ID <= 0 {
			status = http.StatusNoContent
		}
		response.Json(post, status, w)
	} else {
		response.Error("", http.StatusNotFound, w)
	}
}

func validatePost(body models.Post) (err error) {
	if len(body.PostedText) <= 0 && len(body.ImageUrl) <= 0 {
		err = errors.New("Should contains a posted text or an image url")
		return
	}
	if body.UserId == 0 {
		err = errors.New("Should contains creator ID")
		return
	}
	return
}