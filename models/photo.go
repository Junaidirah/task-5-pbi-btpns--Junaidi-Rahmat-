package models

import (
	"time"
)

type (
	Photo struct {
		ID        int       `json:"id"`
		UserID    int       `json:"user_id"`
		Title     string    `json:"title"`
		Caption   string    `json:"caption"`
		PhotoURL  string    `json:"photo_url"`
		Path      string    `json:"path"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
