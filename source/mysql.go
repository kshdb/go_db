package source

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*
mysql数据库对象
*/
type DbMysql struct {
	Uid    string
	Pwd    string
	IP     string
	Port   int
	DbName string
}

/*
数据源
*/
func (d *DbMysql) db() (*sql.DB, error) {
	// 数据源语法："用户名:密码@[连接方式](主机名:端口号)/数据库名"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", d.Uid, d.Pwd, d.IP, d.Port, d.DbName)
	db, err := sql.Open("mysql", dsn) // open() 方法不会真正的与数据库建立连接，只是设置连接需要的参数
	if err != nil {
		//panic(err)
		fmt.Println(err.Error())
	}
	return db, err
}
func (d *DbMysql) Query(_sql string) (*sql.Rows, error) {
	db, _err := d.db()
	defer db.Close()
	rows, _err := db.Query(_sql)
	return rows, _err
}
