package api

import (
	"database/sql"
	"net/http"

	"github.com/FilledEther20/Reg_Bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var newRequest createAccountRequest
	if err := ctx.ShouldBindBodyWithJSON(&newRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := sqlc.CreateAccountParams{
		Owner:    newRequest.Owner,
		Currency: newRequest.Currency,
		Balance:  0,
	}
	acc, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
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
