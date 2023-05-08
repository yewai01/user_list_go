package models

import "time"

type User struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `json:"name"`
	Birthday  int    `json:"birthday"`
	Gender    string `json:"gender"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
