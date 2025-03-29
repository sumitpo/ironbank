package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ DispModel = (*customDispModel)(nil)

type (
	// DispModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDispModel.
	DispModel interface {
		dispModel
		withSession(session sqlx.Session) DispModel
	}

	customDispModel struct {
		*defaultDispModel
	}
)

// NewDispModel returns a model for the database table.
func NewDispModel(conn sqlx.SqlConn) DispModel {
	return &customDispModel{
		defaultDispModel: newDispModel(conn),
	}
}

func (m *customDispModel) withSession(session sqlx.Session) DispModel {
	return NewDispModel(sqlx.NewSqlConnFromSession(session))
}
