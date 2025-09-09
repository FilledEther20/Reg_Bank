package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/FilledEther20/Reg_Bank/db/sqlc"
	"github.com/FilledEther20/Reg_Bank/token"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createAccountRequest struct {
	Currency string `json:"currency" binding:"required,currency"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var newRequest createAccountRequest
	if err := ctx.ShouldBindBodyWithJSON(&newRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := sqlc.CreateAccountParams{
		Owner:    authPayload.Username,
		Currency: newRequest.Currency,
		Balance:  0,
	}
	acc, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, acc)
}

// Get account using ID.
type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var newRequest getAccountRequest
	if err := ctx.BindUri(&newRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	acc, err := server.store.GetAccount(ctx, newRequest.ID)
	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.Username != acc.Owner {
		err := errors.New("account doesnot belong to authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, acc)
}

// Lists account abiding pagination.
type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAccount(ctx *gin.Context) {
	var newRequest listAccountRequest

	if err := ctx.ShouldBindUri(&newRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := sqlc.ListAccountsParams{
		Limit:  newRequest.PageSize,
		Offset: (newRequest.PageID - 1) * newRequest.PageSize,
	}
	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
