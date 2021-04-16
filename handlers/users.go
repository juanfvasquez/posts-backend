package handlers

import (
	"net/http"
	"strconv"
	"regexp"
	"errors"

	"github.com/gorilla/mux"

	"../db"
	"../io/response"
	"../io/request"
	"../models"
	"../crypto"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := db.GetUsers()
	status := http.StatusOK
	if len(users) <= 0 {
		status = http.StatusNoContent
	} 
	response.Json(users, status, w)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if id, ok := vars["id"]; ok {
		userId, err := strconv.Atoi(id)
		if err != nil {
			response.Error("Given ID is not allowed", http.StatusBadRequest, w)
			return
		}
		user := db.GetUser(uint(userId))
		status := http.StatusOK
		if user.ID == 0 {
			status = http.StatusNoContent
		}
		response.Json(user, status, w)
	} else {
		response.Error("", http.StatusNotFound, w)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var body models.User
	request.Json(r, &body)
	err := validateUser(body)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	existingUser := db.GetUserByEmail(body.Email)
	if existingUser.ID > 0 {
		response.Error("Duplicated Email", http.StatusBadRequest, w)
		return
	}
	pass, err := crypto.EncryptText(body.Password)
	if err != nil {
		response.Error("Error encrypting password", http.StatusBadRequest, w)
		return
	}
	body.Password = pass
	user := db.CreateUser(body)
	response.Json(user, http.StatusCreated, w)
}

func validateUser(user models.User) (err error) {
	if len(user.Name) < 5 {
		err = errors.New("Name property must have 5 characters at least")
		return
	}
	regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(user.Email) < 6 || len(user.Email) > 255 || !regex.MatchString(user.Email) {
		err = errors.New("Email format error")
		return
	}
	if len(user.Password) < 6 {
		err = errors.New("Password must have 6 characters at least")
		return
	}
	return
}