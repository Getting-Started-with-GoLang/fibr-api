package models

type Feedback struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}
