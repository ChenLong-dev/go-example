package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"testserver/internal/config"
	"testserver/model"
)

type ServiceContext struct {
	Config config.Config
	Conn   sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := model.NewMysql(c.MysqlConfig)
	return &ServiceContext{
		Config: c,
		Conn:   sqlConn,
	}
}
