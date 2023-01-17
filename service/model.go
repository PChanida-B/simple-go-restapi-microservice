package service

import (
	"time"
)

type Request struct {
	ID        uint   `gorm:"primary_key;auto_increment;not_null"`
	UserName  string `json:"userName"`
	Phone     string `json:"phone"`
	FirstName string `json:"firstName"`
	Lastname  string `json:"lastname"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Response struct {
	ID       uint   `json:"id"`
	UserName string `json:"userName"`
}

type Error struct {
	Error string `json:"error"`
}
