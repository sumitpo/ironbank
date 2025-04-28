package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID            string    `json:"id"            gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	FromAccountID string    `json:"fromAccountId" gorm:"type:uuid;not null"`
	ToAccountID   string    `json:"toAccountId"   gorm:"type:uuid;not null"`
	Amount        float64   `json:"amount"        gorm:"type:decimal(12,2);not null"`
	Memo          string    `json:"memo"          gorm:"type:varchar(255)"`
	CreatedAt     time.Time `json:"createdAt"     gorm:"autoCreateTime"`
	FromAccount   Account   `json:"-"             gorm:"foreignKey:FromAccountID"`
	ToAccount     Account   `json:"-"             gorm:"foreignKey:ToAccountID"`
}

// ProcessTransfer now explicitly requires the DB transaction
func ProcessTransfer(tx *gorm.DB, transaction *Transaction) error {
	// Verify from account has sufficient balance (again in transaction)
	var fromAccount Account
	if err := tx.First(&fromAccount, "id = ?", transaction.FromAccountID).Error; err != nil {
		return err
	}

	if fromAccount.Balance < transaction.Amount {
		return errors.New("insufficient funds")
	}

	// Deduct from source account
	if err := tx.Model(&Account{}).
		Where("id = ?", transaction.FromAccountID).
		Update("balance", gorm.Expr("balance - ?", transaction.Amount)).
		Error; err != nil {
		return err
	}

	// Add to destination account
	if err := tx.Model(&Account{}).
		Where("id = ?", transaction.ToAccountID).
		Update("balance", gorm.Expr("balance + ?", transaction.Amount)).
		Error; err != nil {
		return err
	}

	// Create transaction record
	if err := tx.Create(transaction).Error; err != nil {
		return err
	}

	return nil
}

// GetTransactionsByAccount returns transaction history
func GetTransactionsByAccount(accountID string) ([]Transaction, error) {
	var transactions []Transaction
	err := db.Where("from_account_id = ? OR to_account_id = ?", accountID, accountID).
		Order("created_at DESC").
		Find(&transactions).
		Error
	return transactions, err
}
