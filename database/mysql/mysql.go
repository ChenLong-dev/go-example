package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	userName  string = "root"
	password  string = "12345678"
	ipAddrees string = "127.0.0.1"
	port      int    = 3306
	dbName    string = "dtm"
	charset   string = "utf8"
)

func connectMysql() *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddrees, port, dbName, charset)
	Db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	return Db
}

func queryData(Db *sqlx.DB) {
	rows, err := Db.Query("select id, gid, url, branch_id, op from trans_branch_op")
	if err != nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		return
	}
	for rows.Next() {
		//定义变量接收查询数据
		var id, gid, url, branch_id, op string

		err := rows.Scan(&id, &gid, &url, &branch_id, &op)
		if err != nil {
			fmt.Printf("get data failed, error:[%v]\n", err)
		}
		fmt.Println(id, gid, url, branch_id, op)
	}

	//关闭结果集（释放连接）
	rows.Close()
}

func MysqlClient() {
	var Db *sqlx.DB = connectMysql()
	defer Db.Close()

	queryData(Db)
}

//运行结果：
//1 2019-07-06 11:45:20 anson 123456 技术部 123456@163.com
//3 2019-07-06 11:45:20 johny 123456 技术部 123456@163.com
//4 2019-07-06 11:45:20 johny 123456 技术部 123456@163.com
