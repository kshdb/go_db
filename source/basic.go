package source

import "database/sql"

/*
sql查询方式
*/
const (
	//查询列表
	GetList = iota + 1
	//查询单条对对象
	GetObject
)

/*
数据库操作接口
*/
type Dbsql interface {
	Query(_sql string) (*sql.Rows, error)
}
