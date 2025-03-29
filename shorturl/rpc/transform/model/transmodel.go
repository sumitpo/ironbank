package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TransModel = (*customTransModel)(nil)

type (
	// TransModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTransModel.
	TransModel interface {
		transModel
		withSession(session sqlx.Session) TransModel
	}

	customTransModel struct {
		*defaultTransModel
	}
)

// NewTransModel returns a model for the database table.
func NewTransModel(conn sqlx.SqlConn) TransModel {
	return &customTransModel{
		defaultTransModel: newTransModel(conn),
	}
}

func (m *customTransModel) withSession(session sqlx.Session) TransModel {
	return NewTransModel(sqlx.NewSqlConnFromSession(session))
}
