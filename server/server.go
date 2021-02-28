package server

import (
	"fmt"

	"github.com/cnaize/oat/database"
	"github.com/gin-gonic/gin"
)

type Server struct {
	r  *gin.Engine
	db *database.DB
}

func NewServer(db *database.DB) *Server {
	server := Server{
		r:  gin.Default(),
		db: db,
	}
	questions := server.r.Group("/questions")
	{
		questions.GET("", server.GetQuestions)
		questions.POST("", server.CreateQuestion)
	}
	return &server
}

func (s *Server) Run(port uint) error {
	fmt.Printf("Server run on port: %d\n", port)
	return s.r.Run(fmt.Sprintf(":%d", port))
}
