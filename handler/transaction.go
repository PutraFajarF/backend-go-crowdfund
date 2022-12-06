package handler

import (
	"go-crowdfunding/helper"
	"go-crowdfunding/transaction"
	"go-crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// parameter di uri
// tangkap parameter, mapping ke input struct
// panggil service, input struct sebagai parameter
// service, berbekal campaign id bisa panggil repo
// repo mencari data trasaction suatu campaign

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionsDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionsByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign transactions", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

// flow GetUserTransactions
// Perlu diketahui siapa user yg login
// handler -> ambil nilai user dari jwt/middleware -> didapatkan userID
// service
// repository -> ambil data transactions dan preload data campaign karena pada JSON yg ingin ditampilkan tdk hanya data transaction
// namun data nama campaign dan image campaign yg terkait dgn Transactions tersebut
