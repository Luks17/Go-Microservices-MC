package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Luks17/Go-Microservices-MC/api/handlers/errmap"
	"github.com/Luks17/Go-Microservices-MC/db"
	"github.com/Luks17/Go-Microservices-MC/db/repository"
	"github.com/Luks17/Go-Microservices-MC/db/sqlc"
	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	FromAccountID int64           `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64           `json:"to_account_id" binding:"required,min=1"`
	Amount        float64         `json:"amount" binding:"required,gt=0"`
	Currency      sqlc.Currencies `json:"currency" binding:"required,currency"`
}

func CreateTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errmap.ErrorResponse(err))
		return
	}

	if !validAccount(ctx, req.FromAccountID, req.Currency) {
		return
	}
	if !validAccount(ctx, req.ToAccountID, req.Currency) {
		return
	}

	amount := strconv.FormatFloat(req.Amount, 'f', 2, 64)
	arg := repository.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        amount,
	}

	result, err := db.DBStore.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errmap.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func validAccount(ctx *gin.Context, accountID int64, currency sqlc.Currencies) bool {
	account, err := db.DBStore.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errmap.ErrorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errmap.ErrorResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", accountID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errmap.ErrorResponse(err))
		return false
	}

	return true
}
