package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DbWorker struct {
	Dsn string
	Db  *sql.DB
}

type Cate struct {
	cid     int
	cname   string
	addtime int
	scope   int
}


//通常我们使用 defer rows.Close() 来确保数据库连接可以正确放回到连接池中。
func main() {
	dbw := DbWorker{Dsn: "root:123456@tcp(localhost:3306)/mydb?charset=utf8mb4"}
	// 支持下面几种DSN写法，具体看mysql服务端配置，常见为第2种
	// user@unix(/path/to/socket)/dbname?charset=utf8
	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	// user:password@/dbname
	// user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname

	dbtemp,  err := sql.Open("mysql",  dbw.Dsn)
	dbw.Db = dbtemp

	if err != nil {
		panic(err)
		return
	}
	defer dbw.Db.Close()

	// 插入数据测试
	dbw.insertData()

	// 删除数据测试
	dbw.deleteData()

	// 修改数据测试
	dbw.editData()

	// 查询数据测试
	dbw.queryData()

	// 事务操作测试
	dbw.transaction()
}

