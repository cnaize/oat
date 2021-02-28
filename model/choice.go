package model

// Choice model
type Choice struct {
	Text string `json:"text" binding:"required"`
}
