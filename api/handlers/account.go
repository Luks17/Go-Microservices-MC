package handlers

import (
	"net/http"

	"github.com/Luks17/Go-Microservices-MC/db"
	"github.com/Luks17/Go-Microservices-MC/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string          `json:"owner" binding:"required"`
	Currency sqlc.Currencies `json:"currency" binding:"required,oneof=USD EUR BRL"`
}

func CreateAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := sqlc.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  "0.00",
		Currency: req.Currency,
	}

	account, err := db.DBStore.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
