package handlers

import (
	"net/http"
	
	"golang.org/x/crypto/bcrypt"

	"postsbackend/db"
	"postsbackend/models"
	"postsbackend/jwt"
	"postsbackend/io/response"
	"postsbackend/io/request"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var body models.User
	err := request.Json(r, &body)
	if err != nil {
		response.Error("Malformed body", http.StatusBadRequest, w)
		return
	}
	if len(body.Email) <= 0 || len(body.Password) <= 0 {
		response.Error("Email and password required", http.StatusBadRequest, w)
		return
	}
	user := db.GetUserByEmail(body.Email)
	if user.ID <= 0 {
		response.Error("User does not exist", http.StatusBadRequest, w)
		return
	}
	userPass := []byte(user.Password)
	bodyPass := []byte(body.Password)
	err = bcrypt.CompareHashAndPassword(userPass, bodyPass)
	if err != nil {
		response.Error("Wrong password", http.StatusBadRequest, w)
		return
	}
	token, err := jwt.GenerateToken(user)
	if err != nil {
		response.Error("Error generating JWT Token", http.StatusInternalServerError, w)
		return
	}
	user.Password = ""
	loginResp := models.LoginResponse{Token: token, User: user}
	response.Json(loginResp, http.StatusOK, w)
}