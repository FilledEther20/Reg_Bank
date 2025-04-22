package api

import (
	"github.com/FilledEther20/Reg_Bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  sqlc.Store
	router *gin.Engine
}

// To create a New http server and setup routing
func NewServer(store sqlc.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount) // For listing accounts in a page by using query params

	server.router = router
	return server
}

// Starts and runs a HTTP server on a given address.
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
