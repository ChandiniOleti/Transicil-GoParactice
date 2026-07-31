package handlers

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"

	"paylaterservice/generated"
	"paylaterservice/services"
)

type TransactionRequest struct {
    MerchantID int32  `json:"merchant_id"`
    Amount     string `json:"amount"`
}

// Create Transaction
func CreateTransaction(c *gin.Context) {

	var request TransactionRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := int32(c.GetInt("user_id"))

	transaction := generated.CreateTransactionParams{
    UserID: userID,
		MerchantID: sql.NullInt32{
			Int32: request.MerchantID,
			Valid: true,
		},
		Amount: request.Amount,
	}

	response, err := services.ProcessTransaction(transaction)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, response)
}

// Get All Transactions
func GetTransactions(c *gin.Context) {

	transactions, err := services.GetTransactions()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, transactions)
}

// Get Transaction By ID
func GetTransactionByID(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := services.GetTransactionByID(int32(id))

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, transaction)
}

// Get Transactions By User
func GetTransactionsByUser(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("user_id"))

	transactions, err := services.GetTransactionsByUser(int32(id))

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, transactions)
}

// Get Transactions By Merchant
func GetTransactionsByMerchant(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("merchant_id"))

	transactions, err := services.GetTransactionsByMerchant(int32(id))

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, transactions)
}