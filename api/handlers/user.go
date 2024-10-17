package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Luks17/Go-Microservices-MC/api/handlers/errmap"
	"github.com/Luks17/Go-Microservices-MC/crypt"
	"github.com/Luks17/Go-Microservices-MC/db"
	"github.com/Luks17/Go-Microservices-MC/db/repository/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanumunicode"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type createUserResponse struct {
	Username              string       `json:"username"`
	FullName              string       `json:"full_name"`
	Email                 string       `json:"email"`
	PasswordLastChangedAt sql.NullTime `json:"password_last_changed_at"`
	CreatedAt             time.Time    `json:"created_at"`
}

func CreateUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errmap.ErrorResponse(err))
		return
	}

	hashedPassword, err := crypt.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errmap.ErrorResponse(err))
		return
	}

	arg := sqlc.CreateUserParams{
		Username: req.Username,
		Password: hashedPassword,
		FullName: req.FullName,
		Email:    req.Email,
	}

	user, err := db.DBStore.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errmap.ErrorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errmap.ErrorResponse(err))
		return
	}

	res := createUserResponse{
		Username:              user.Username,
		FullName:              user.FullName,
		Email:                 user.Email,
		CreatedAt:             user.CreatedAt,
		PasswordLastChangedAt: user.PasswordLastChangedAt,
	}
	ctx.JSON(http.StatusOK, res)
}
