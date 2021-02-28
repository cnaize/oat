package model

import (
	"time"
)

// Question model
type Question struct {
	Text string `json:"text" binding:"required"`
	// Choices associated to the question
	Choices []Choice `json:"choices" binding:"required"`
	// Creation date of the question
	CreatedAt time.Time `json:"createdAt"`
}
