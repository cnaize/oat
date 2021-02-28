package server

import (
	"net/http"

	"github.com/cnaize/oat/model"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetQuestions(c *gin.Context) {
	questions, err := s.db.Questions.List()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, questions)
}

func (s *Server) CreateQuestion(c *gin.Context) {
	var question model.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})
		return
	}
	if err := s.db.Questions.Create(question); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
