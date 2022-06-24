package model

import (
	"time"
)

type OneResponse struct {
	Id          int       `json:"id"`
	Category    int       `json:"category,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Footer      string    `json:"footer,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}

type ThreeResponse struct {
	Id        int                 `json:"id"`
	Category  int                 `json:"category,omitempty"`
	Items     []ThreeItemResponse `json:"items,omitempty"`
	Tags      []string            `json:"tags,omitempty"`
	CreatedAt time.Time           `json:"createdAt,omitempty"`
}

type ThreeItemResponse struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Footer      string `json:"footer,omitempty"`
}
