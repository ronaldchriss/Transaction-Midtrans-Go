package handler

import (
	"bwa_go/helper"
	"bwa_go/transaction"
	"bwa_go/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *TransactionHandler {
	return &TransactionHandler{service}
}

func (h *TransactionHandler) GetTransaction(c *gin.Context) {
	var input transaction.InputGetTransaction

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.JsonResponse("Error to Get Transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	code := c.MustGet("codeUser").(user.User)
	input.User = code

	trans, err := h.service.GetTransaction(input)
	if err != nil {
		response := helper.JsonResponse("Error to Get Transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JsonResponse("Transaction Detail", http.StatusOK, "success", transaction.FormatListTrans(trans))
	c.JSON(http.StatusOK, response)
}
