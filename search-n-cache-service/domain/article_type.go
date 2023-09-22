package domain

import (
	"time"
)

type ArticleType struct {
	ID              int64     `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Author          string    `json:"author"`
	ArticleCategory string    `json:"category"`
	ArticleType     string    `json:"type"`
	Content         string    `json:"content"`
	Tags            string    `json:"tags"`
	CreatedAt       time.Time `json:"createdAt"`
	CreatedBy       string    `json:"createdBy"`
	UpdatedAt       time.Time `json:"updatedAt"`
	UpdatedBy       string    `json:"updatedBy"`
}
