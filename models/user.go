package models

import (
	"time"
)

type (
	User struct {
		ID        int       `json:"id"`
		UserName  string    `json:"name"`
		Email     string    `json:"Email"`
		Password  string    `json:"Password"`
		Photos    []Photo   `json:"photos" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Relasi dengan Photo
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)



