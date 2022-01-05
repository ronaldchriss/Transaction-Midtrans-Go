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

func (h *TransactionHandler) GetTransactionByUser(c *gin.Context) {
	code := c.MustGet("codeUser").(user.User)
	UserID := code.ID

	trans, err := h.service.GetByUserID(UserID)
	if err != nil {
		response := helper.JsonResponse("Error to Get Transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.JsonResponse("User's Transaction ", http.StatusOK, "success", transaction.FormatListTransUser(trans))
	c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) CreateTrans(c *gin.Context) {
	var input transaction.InputCreateTrans

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMassage := gin.H{"errors": errors}

		response := helper.JsonResponse("Save Transaction Failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	code := c.MustGet("codeUser").(user.User)

	input.User = code

	trans, err := h.service.CreateTrans(input)
	if err != nil {
		response := helper.JsonResponse("Save Transaction Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := transaction.SuccessFormatter(trans)
	response := helper.JsonResponse("Save Transaction Success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
