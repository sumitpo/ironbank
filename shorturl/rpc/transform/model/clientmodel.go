package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ClientModel = (*customClientModel)(nil)

type (
	// ClientModel is an interface to be customized, add more methods here,
	// and implement the added methods in customClientModel.
	ClientModel interface {
		clientModel
		withSession(session sqlx.Session) ClientModel
	}

	customClientModel struct {
		*defaultClientModel
	}
)

// NewClientModel returns a model for the database table.
func NewClientModel(conn sqlx.SqlConn) ClientModel {
	return &customClientModel{
		defaultClientModel: newClientModel(conn),
	}
}

func (m *customClientModel) withSession(session sqlx.Session) ClientModel {
	return NewClientModel(sqlx.NewSqlConnFromSession(session))
}
