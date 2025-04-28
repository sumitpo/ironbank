package models

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type AccountType string

const (
	Checking AccountType = "checking"
	Savings  AccountType = "savings"
)

type Account struct {
	ID        string      `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID    string      `json:"userId" gorm:"type:uuid;not null"`
	Number    string      `json:"number" gorm:"unique;not null"`
	Type      AccountType `json:"type" gorm:"type:varchar(20);not null"`
	Balance   float64     `json:"balance" gorm:"type:decimal(12,2);not null;default:0"`
	CreatedAt time.Time   `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time   `json:"updatedAt" gorm:"autoUpdateTime"`
	User      User        `json:"-" gorm:"foreignKey:UserID"`
}

// GenerateAccountNumber creates a random bank account number
func GenerateAccountNumber() string {
	var sb strings.Builder
	for i := 0; i < 3; i++ {
		sb.WriteString(fmt.Sprintf("%04d", rand.Intn(10000)))
		if i < 2 {
			sb.WriteString("-")
		}
	}
	return sb.String()
}

// CreateAccount creates a new bank account
func CreateAccount(account *Account) error {
	if err := db.Create(account).Error; err != nil {
		return err
	}
	return nil
}

// GetAccountsByUserID returns all accounts for a user
func GetAccountsByUserID(userID string) ([]Account, error) {
	var accounts []Account
	if err := db.Where("user_id = ?", userID).Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

// GetAccountByID finds an account by its ID
func GetAccountByID(id string) (*Account, error) {
	var account Account
	if err := db.First(&account, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

// VerifyAccountOwnership checks if an account belongs to a user
func VerifyAccountOwnership(accountID, userID string) (*Account, error) {
	var account Account
	if err := db.First(&account, "id = ? AND user_id = ?", accountID, userID).Error; err != nil {
		return nil, err
	}
	return &account, nil
}
