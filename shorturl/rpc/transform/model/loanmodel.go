package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ LoanModel = (*customLoanModel)(nil)

type (
	// LoanModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLoanModel.
	LoanModel interface {
		loanModel
		withSession(session sqlx.Session) LoanModel
	}

	customLoanModel struct {
		*defaultLoanModel
	}
)

// NewLoanModel returns a model for the database table.
func NewLoanModel(conn sqlx.SqlConn) LoanModel {
	return &customLoanModel{
		defaultLoanModel: newLoanModel(conn),
	}
}

func (m *customLoanModel) withSession(session sqlx.Session) LoanModel {
	return NewLoanModel(sqlx.NewSqlConnFromSession(session))
}
