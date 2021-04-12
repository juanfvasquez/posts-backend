package models

type Message struct {
	Event 	string 	`json:"event"`
	Post 		Post 		`json:"post,omitempty"`
	PostId 	uint 		`json:"post_id,omitempty"`
	Comment Comment `json:"comment,omitempty"`
}