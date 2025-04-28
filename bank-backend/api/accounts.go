package api

import (
	"net/http"

	"bank-backend/models"

	"github.com/gin-gonic/gin"
)

type CreateAccountRequest struct {
	Type    models.AccountType `json:"type"    binding:"required,oneof=checking savings"`
	Balance float64            `json:"balance" binding:"min=0"`
}

// GetAccounts returns all accounts for the authenticated user
func GetAccounts(c *gin.Context) {
	userID := c.MustGet("userID").(string)

	accounts, err := models.GetAccountsByUserID(userID)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to fetch accounts"},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accounts": accounts,
	})
}

// CreateAccount creates a new bank account
func CreateAccount(c *gin.Context) {
	userID := c.MustGet("userID").(string)

	var req CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account := models.Account{
		UserID:  userID,
		Type:    req.Type,
		Balance: req.Balance,
		Number:  models.GenerateAccountNumber(),
	}

	if err := models.CreateAccount(&account); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to create account"},
		)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Account created successfully",
		"account": account,
	})
}
