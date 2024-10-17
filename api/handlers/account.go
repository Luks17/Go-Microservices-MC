package handlers

import (
	"database/sql"
	"net/http"

	"github.com/Luks17/Go-Microservices-MC/api/handlers/errmap"
	"github.com/Luks17/Go-Microservices-MC/db"
	"github.com/Luks17/Go-Microservices-MC/db/repository/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type listAccountsRequest struct {
	PageID   int32 `form:"page_id,default=1" binding:"min=1"`
	PageSize int32 `form:"page_size,default=5" binding:"min=5"`
}

func ListAccounts(ctx *gin.Context) {
	var req listAccountsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errmap.ErrorResponse(err))
		return
	}

	arg := sqlc.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := db.DBStore.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errmap.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func GetAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errmap.ErrorResponse(err))
		return
	}

	account, err := db.DBStore.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errmap.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errmap.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type createAccountRequest struct {
	Owner    string          `json:"owner" binding:"required"`
	Currency sqlc.Currencies `json:"currency" binding:"required,currency"`
}

func CreateAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errmap.ErrorResponse(err))
		return
	}

	arg := sqlc.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  "0.00",
		Currency: req.Currency,
	}

	account, err := db.DBStore.CreateAccount(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errmap.ErrorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errmap.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
