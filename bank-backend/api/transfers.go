package api

import (
	"net/http"

	"bank-backend/models"

	"github.com/gin-gonic/gin"
)

type TransferRequest struct {
	FromAccountID string  `json:"fromAccountId" binding:"required,uuid"`
	ToAccountID   string  `json:"toAccountId"   binding:"required,uuid"`
	Amount        float64 `json:"amount"        binding:"required,gt=0"`
	Memo          string  `json:"memo"          binding:"max=255"`
}

func TransferFunds(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	db := models.GetDB()

	var req TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify the fromAccount belongs to the user
	fromAccount, err := models.GetAccountByID(req.FromAccountID)
	if err != nil || fromAccount.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid source account"})
		return
	}

	// Verify sufficient balance
	if fromAccount.Balance < req.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient funds"})
		return
	}

	// Verify destination account exists
	_, err = models.GetAccountByID(req.ToAccountID)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Invalid destination account"},
		)
		return
	}

	// Create transaction record
	transaction := models.Transaction{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
		Memo:          req.Memo,
	}

	// Get a database transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Process the transfer within the transaction
	if err := models.ProcessTransfer(tx, &transaction); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to commit transaction"},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Transfer completed successfully",
		"transaction": transaction,
	})
}
