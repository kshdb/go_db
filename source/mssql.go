package source

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

/*
mssql数据库对象
*/
type DbMssql struct {
	Uid    string
	Pwd    string
	IP     string
	Port   int
	DbName string
}

/*
数据源
*/
func (d *DbMssql) db() (*sql.DB, error) {
	// 数据源语法："用户名:密码@[连接方式](主机名:端口号)/数据库名"
	dsn := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%d", d.IP, d.DbName, d.Uid, d.Pwd, d.Port)
	db, err := sql.Open("mssql", dsn) // open() 方法不会真正的与数据库建立连接，只是设置连接需要的参数
	if err != nil {
		//panic(err)
		fmt.Println(err.Error())
	}
	return db, err
}
func (d *DbMssql) Query(_sql string) (*sql.Rows, error) {
	db, _err := d.db()
	defer db.Close()
	rows, _err := db.Query(_sql)
	return rows, _err
}
