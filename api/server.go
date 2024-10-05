package api

import (
	"github.com/Luks17/Go-Microservices-MC/api/handlers"
	"github.com/gin-gonic/gin"
)

func InitServer(address string) error {
	server := gin.Default()

	server.POST("/accounts", handlers.CreateAccount)

	err := server.Run(address)
	if err != nil {
		return err
	}

	return nil
}
