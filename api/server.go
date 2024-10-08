package api

import (
	"github.com/Luks17/Go-Microservices-MC/api/handlers"
	"github.com/Luks17/Go-Microservices-MC/api/validators"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func LoadValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validators.ValidCurrency)
	}
}

func LoadRouter(server *gin.Engine) {
	server.GET("/accounts", handlers.ListAccounts)
	server.GET("/accounts/:id", handlers.GetAccount)
	server.POST("/accounts", handlers.CreateAccount)

	server.POST("/transfers", handlers.CreateTransfer)
}

func InitServer(address string) error {
	server := gin.Default()

	LoadValidators()
	LoadRouter(server)

	err := server.Run(address)
	if err != nil {
		return err
	}

	return nil
}
