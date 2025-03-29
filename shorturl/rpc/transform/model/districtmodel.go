package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ DistrictModel = (*customDistrictModel)(nil)

type (
	// DistrictModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDistrictModel.
	DistrictModel interface {
		districtModel
		withSession(session sqlx.Session) DistrictModel
	}

	customDistrictModel struct {
		*defaultDistrictModel
	}
)

// NewDistrictModel returns a model for the database table.
func NewDistrictModel(conn sqlx.SqlConn) DistrictModel {
	return &customDistrictModel{
		defaultDistrictModel: newDistrictModel(conn),
	}
}

func (m *customDistrictModel) withSession(session sqlx.Session) DistrictModel {
	return NewDistrictModel(sqlx.NewSqlConnFromSession(session))
}
