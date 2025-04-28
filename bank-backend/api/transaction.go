package api

import (
	"net/http"

	"bank-backend/models"

	"github.com/gin-gonic/gin"
)

// GetTransactions returns transaction history for an account
func GetTransactions(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	accountID := c.Param("accountId")

	// Verify account belongs to user
	if _, err := models.VerifyAccountOwnership(accountID, userID); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Account access denied"})
		return
	}

	transactions, err := models.GetTransactionsByAccount(accountID)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to fetch transactions"},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
	})
}
