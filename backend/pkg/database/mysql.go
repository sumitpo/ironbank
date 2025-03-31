package database

import (
	"bank-app/pkg/models"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	db *sql.DB
}

func NewMySQLConnection(connectionString string) (*MySQLDB, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &MySQLDB{db: db}, nil
}

func (m *MySQLDB) AuthenticateUser(username, password string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, full_name FROM users WHERE username = ? AND password = SHA2(?, 256)`

	err := m.db.QueryRow(query, username, password).Scan(
		&user.ID,
		&user.Username,
		&user.FullName,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *MySQLDB) GetAccounts(userID int) ([]models.Account, error) {
	query := `SELECT id, user_id, account_number, balance, type, created_at
	          FROM accounts WHERE user_id = ?`

	rows, err := m.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []models.Account
	for rows.Next() {
		var acc models.Account
		err := rows.Scan(
			&acc.ID,
			&acc.UserID,
			&acc.AccountNumber,
			&acc.Balance,
			&acc.Type,
			&acc.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, acc)
	}

	return accounts, nil
}
