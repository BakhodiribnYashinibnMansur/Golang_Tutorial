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

type ArticleRequest struct {
	ID      int    `json:"id"`
	Content        // Promoted fiels
	Author  Person `json:"author"`
}

type ArticleCreate struct {
	Content        // Promoted fiels
	Author  Person `json:"author"`
}

type Query struct {
	Offset int    `json:"offset" default:"0"`
	Limit  int    `json:"limit" default:"10"`
	Search string `json:"search"`
}
