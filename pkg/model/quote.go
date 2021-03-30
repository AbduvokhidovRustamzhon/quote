package model

import (
	"time"
)


type Quote struct {
	ID        string    `json:"id"`
	Author    string    `json:"author"`
	Quote     string    `json:"quote"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}

