package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Claim struct {
	ID									uint		`json:"id"`
	Name 								string	`json:"name"`
	Email 							string 	`json:"email"`
	jwt.StandardClaims
}