package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"testserver/internal/config"
	"time"
)

func NewMysql(mysqlConfig config.MysqlConfig) sqlx.SqlConn {
	mysql := sqlx.NewMysql(mysqlConfig.DataSource)
	db, err := mysql.RawDB()
	if err != nil {
		panic(err)
	}
	cxt, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(mysqlConfig.ConnectTimeout))
	defer cancel()
	err = db.PingContext(cxt)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return mysql
}
