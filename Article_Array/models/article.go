package models

import (
	"time"
)

type Content struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Article struct {
	ID        int        `json:"id"`
	Content              // Promoted fields
	Author    Person     `json:"author"` // Nested struct
	CreatedAt *time.Time `json:"created_at"`
}
