package main

import (
	"fmt"

	"github.com/kshdb/go_db/source"
)

func main() {
	//mysql操作示例
	_se_mysql := source.SqlEntry{
		Type:      source.MySql,
		QueryType: source.GetList,
		Auth: source.DbAuth{
			Uid:    "root",
			Pwd:    "haosql",
			IP:     "127.0.0.1",
			Port:   3306,
			DbName: "test2",
		},
		QuerySql: "select id,title,path from post limit 10",
	}
	//mysql list查询
	_json, _ := _se_mysql.GetJson()
	fmt.Printf("当前mysql---json是%s\n", _json)

}
