package handlers

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"

	"transactionservice/generated"
	"transactionservice/services"
)

type TransactionRequest struct {
	MerchantID int32  `json:"merchant_id"`
	Amount     string `json:"amount"`
}

// CreateTransaction creates a new purchase transaction.
// POST /transactions
func CreateTransaction(c *gin.Context) {
	var request TransactionRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID := int32(c.GetInt("user_id"))
	authHeader := c.GetHeader("Authorization")

	transaction := generated.CreateTransactionParams{
		UserID: userID,
		MerchantID: sql.NullInt32{
			Int32: request.MerchantID,
			Valid: true,
		},
		Amount: request.Amount,
	}

	response, err := services.ProcessTransaction(transaction, authHeader)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, response)
}

// GetTransactions returns all transactions.
// GET /transactions
func GetTransactions(c *gin.Context) {
	transactions, err := services.GetTransactions()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, transactions)
}

// GetTransactionByID returns a single transaction.
// GET /transactions/:id
func GetTransactionByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := services.GetTransactionByID(int32(id))
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, transaction)
}

// GetTransactionsByUser returns all transactions for a user.
// GET /transactions/user/:user_id
func GetTransactionsByUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid User ID"})
		return
	}

	transactions, err := services.GetTransactionsByUser(int32(id))
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, transactions)
}

// GetTransactionsByMerchant returns all transactions for a merchant.
// GET /transactions/merchant/:merchant_id
func GetTransactionsByMerchant(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("merchant_id"))

	transactions, err := services.GetTransactionsByMerchant(int32(id))
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, transactions)
}

// PaybackTransactionRequest is the body for internal PAYBACK creation.
type PaybackTransactionRequest struct {
	UserID int32  `json:"user_id"`
	Amount string `json:"amount"`
}

// CreatePaybackTransaction records a PAYBACK transaction (internal only).
// POST /internal/transactions/payback
func CreatePaybackTransaction(c *gin.Context) {
	var request PaybackTransactionRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response, err := services.CreatePaybackTransaction(request.UserID, request.Amount)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, response)
}

// GetTransactionsInternal returns all transactions (internal only).
// GET /internal/transactions
func GetTransactionsInternal(c *gin.Context) {
	transactions, err := services.GetTransactions()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, transactions)
}

// GetTransactionsByUserInternal returns transactions for a user (internal only).
// GET /internal/transactions/user/:user_id
func GetTransactionsByUserInternal(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid User ID"})
		return
	}

	transactions, err := services.GetTransactionsByUser(int32(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, transactions)
}

// GetTransactionsByMerchantInternal returns transactions for a merchant (internal only).
// GET /internal/transactions/merchant/:merchant_id
func GetTransactionsByMerchantInternal(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Merchant ID"})
		return
	}

	transactions, err := services.GetTransactionsByMerchant(int32(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, transactions)
}
