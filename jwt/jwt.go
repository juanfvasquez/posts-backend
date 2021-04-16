package jwt

import (
	"errors"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"../models"
)

var JWT_KEY = []byte("softlond.co@golang_course_2021")

// UserEmail keep user email
var UserEmail string

// UserID keep user ID
var UserID string

// GenerateToken returns jwt token or error
func GenerateToken(user models.User) (string, error) {
	payload := jwt.MapClaims{
		"email"	: user.Email,
		"name"	:  user.Name,
		"id"		:  user.ID,
		"exp"		: time.Now().Add(time.Hour * 24).Unix(),
		// email, roles...
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString(JWT_KEY)
	return token, err
}

// ProcessToken check if token is valid
func ProcessToken(authHeader string) (*models.Claim, error) {
	claims := &models.Claim{}
	if len(authHeader) <= 0 {
		return claims, errors.New("No auth token present")
	}
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return claims, errors.New("Invalid token format")
	}
	tokenStr := strings.TrimSpace(splitToken[1])
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JWT_KEY, nil
	})
	if err != nil {
		return claims, err
	}
	// exists := db.GetUserByEmail(claims.Email)
	// if exists.ID >= 0 {
	// 	UserEmail = claims.Email
	// 	UserID = ID
	// }
	return claims, nil
}
