package source

import (
	"encoding/json"
	"fmt"
)

/*
sql数据库类型
*/
const (
	MySql = iota + 101
	MSSql
	Oracle
	Sqlite
	Access
	Excel
	Txt
	Csv
)

/*
sql数据库操作对象
*/
type SqlEntry struct {
	Type      int    `json:"type"`
	Auth      DbAuth `json:"auth"`
	QueryType int    `json:"query_type"`
	QuerySql  string `json:"query_sql"`
}

/*
数据库鉴权
*/
type DbAuth struct {
	Uid    string
	Pwd    string
	IP     string
	Port   int
	DbName string //如果是文件数据库这里填写的是文件地址
}

/*
sql数据源
*/
func (s *SqlEntry) GetJson() (_rt_json string, err error) {
	//_rt_json := ""
	switch s.QueryType {
	case GetList:
		_list, _err := s.get_list(s.QuerySql)
		err = _err
		_json, _err1 := json.Marshal(_list)
		err = _err1
		_rt_json = string(_json)
	case GetObject:
		_model, _err := s.get_one(s.QuerySql)
		err = _err
		_json, _err1 := json.Marshal(_model)
		err = _err1
		_rt_json = string(_json)
	}
	return
}

/*
列表查询
*/
func (s *SqlEntry) get_list(_sql string) (tableData []map[string]interface{}, err error) {
	//tableData := make([]map[string]interface{}, 0)
	var _db Dbsql
	switch s.Type {
	case MySql:
		_db = &DbMysql{Uid: s.Auth.Uid, Pwd: s.Auth.Pwd, IP: s.Auth.IP, Port: s.Auth.Port, DbName: s.Auth.DbName}
	case MSSql:
		_db = &DbMssql{Uid: s.Auth.Uid, Pwd: s.Auth.Pwd, IP: s.Auth.IP, Port: s.Auth.Port, DbName: s.Auth.DbName}
	case Sqlite:
		_db = &DbSqlite{Pwd: s.Auth.Pwd, DbName: s.Auth.DbName}
	case Access:
		_db = &DbAccess{Pwd: s.Auth.Pwd, DbName: s.Auth.DbName}
	}
	rows, _err := _db.Query(_sql)
	if _err != nil {
		err = _err
		fmt.Println("数据读取出错了", _err)
	}
	defer rows.Close()
	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	scanArgs := make([]interface{}, count)
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		rows.Scan(scanArgs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			v := values[i]
			b, ok := v.([]byte)
			if ok {
				entry[col] = string(b)
			} else {
				entry[col] = v
			}
		}
		tableData = append(tableData, entry)
	}
	return
}

/*
单条查询
*/
func (s *SqlEntry) get_one(_sql string) (_one map[string]interface{}, err error) {
	//_one := map[string]interface{}{}
	var _db Dbsql
	switch s.Type {
	case MySql:
		_db = &DbMysql{Uid: s.Auth.Uid, Pwd: s.Auth.Pwd, IP: s.Auth.IP, Port: s.Auth.Port, DbName: s.Auth.DbName}
	case MSSql:
		_db = &DbMssql{Uid: s.Auth.Uid, Pwd: s.Auth.Pwd, IP: s.Auth.IP, Port: s.Auth.Port, DbName: s.Auth.DbName}
	case Sqlite:
		_db = &DbSqlite{Pwd: s.Auth.Pwd, DbName: s.Auth.DbName}
	case Access:
		_db = &DbAccess{Pwd: s.Auth.Pwd, DbName: s.Auth.DbName}
	}
	rows, _ := _db.Query(_sql)
	defer rows.Close()
	columns, _err1 := rows.Columns()
	if _err1 != nil {
		err = _err1
		fmt.Println("数据读取出错了", _err1)
	}
	count := len(columns)
	values := make([]interface{}, count)
	scanArgs := make([]interface{}, count)
	for i := range values {
		scanArgs[i] = &values[i]
	}
	if rows.Next() {
		rows.Scan(scanArgs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			v := values[i]
			b, ok := v.([]byte)
			if ok {
				entry[col] = string(b)
			} else {
				entry[col] = v
			}
		}
		_one = entry
	}
	return
}
