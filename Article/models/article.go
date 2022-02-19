package models

import (
	"time"
)

type Content struct {
	Title string
	Body  string
}

type Article struct {
	ID        int
	Content          // Promoted fields
	Author    Person // Nested struct
	CreatedAt *time.Time
}
