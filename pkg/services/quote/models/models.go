package models

import (
	"github.com/google/uuid"
	"time"
)

type Quote struct {
	ID        string    `json:"id"`
	Author    string    `json:"author"`
	Quote     string    `json:"quote"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}

type Request struct {
	Author   string `json:"author"`
	Quote    string `json:"quote"`
	Category string `json:"category"`
}

func NewQuote(author string, quote string, category string) *Quote {
	return &Quote{
		ID:        uuid.New().String(),
		Author:    author,
		Quote:     quote,
		Category:  category,
		CreatedAt: time.Now(),
	}
}
