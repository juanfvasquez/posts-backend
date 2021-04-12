package models

type LoginResponse struct {
	Token string `json:"token,omitempty"`
	User User `json:"user"`
}