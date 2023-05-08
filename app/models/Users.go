package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Birthday int    `json:"birthday"`
	Gender   string `json:"gender"`
}
