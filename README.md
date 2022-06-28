# github.com/kshdb/go_db

go 原生数据库操作

# 操作示例

```
import (
	"fmt"

	"github.com/kshdb/go_db/source"
)

func main() {
	/*_semssql := source.SqlEntry{
		Type:      source.MSSql,
		QueryType: source.GetList,
		Auth: source.DbAuth{
			Uid:    "sa",
			Pwd:    "sa",
			IP:     "127.0.0.1",
			Port:   1433,
			DbName: "JHdata_2016_Data",
		},
		QuerySql: "select top 10 cczz,bbzz from cizu",
	}
	//mysql list查询
	_jsonmssql := _semssql.GetJson()
	fmt.Printf("当前mssql---_sonmssql是%s\n", _jsonmssql)
	_semssql.QueryType = source.GetObject
	_semssql.QuerySql = "select top 1 cczz,bbzz from cizu"
	_jsonmssql1 := _semssql.GetJson()
	fmt.Printf("当前mssql---sjsonmssql1是%s\n", _jsonmssql1)
	//_semssql.QueryType=source.GetObject
	_semssql.QuerySql = "select  count(by1) as aaa, sum(gls) as bbb from cizu"
	_jsonmssql2 := _semssql.GetJson()
	fmt.Printf("当前mssql---sjsonmssql2是%s\n", _jsonmssql2)
	*/
	//mysql操作示例
	_se_mysql := source.SqlEntry{
		Type:      source.MySql,
		QueryType: source.GetList,
		Auth: source.DbAuth{
			Uid:    "root",
			Pwd:    "haosql",
			IP:     "127.0.0.1",
			Port:   3306,
			DbName: "blog",
		},
		QuerySql: "select id,title,path from post limit 10",
	}
	//mysql list查询
	_json, _ := _se_mysql.GetJson()
	fmt.Printf("当前mysql---json是%s\n", _json)
	_se_mysql.QueryType = source.GetObject
	_se_mysql.QuerySql = "select title,path from post where id=13 limit 1"
	_json1, _ := _se_mysql.GetJson()
	fmt.Printf("当前mysql---json1是%s\n", _json1)
	_se_mysql.QuerySql = "select count(id) as aaa,sum(id) as bbb from post"
	_json2, _ := _se_mysql.GetJson()
	fmt.Printf("当前mysql---json2是%s\n", _json2)
	//source.DbNoSql(source.Mongo)

	//sqlite操作示例
	_se_sqlite := source.SqlEntry{
		Type:      source.Sqlite,
		QueryType: source.GetList,
		Auth: source.DbAuth{
			Pwd:    "888888",
			DbName: "./test1.db",
		},
		QuerySql: "SELECT id,name FROM users ORDER BY id",
	}
	//sqlite list查询
	_json_sqlite, _ := _se_sqlite.GetJson()
	fmt.Printf("当前sqlite---json是%s\n", _json_sqlite)

	//access操作示例
	_se_access := source.SqlEntry{
		Type:      source.Access,
		QueryType: source.GetList,
		Auth: source.DbAuth{
			Pwd:    "123456",
			DbName: "./db1.mdb",
		},
		QuerySql: "SELECT * from InInfo",
	}
	//access list查询
	_json_access, _ := _se_access.GetJson()
	fmt.Printf("当前access---json是%s\n", _json_access)

	//excel操作示例
	_se_excel := source.DocEntry{
		Type:      source.Excel,
		QueryType: source.GetList,
		Auth: source.DbAuth{
			DbName: "./test.xlsx",
		},
	}
	//excel list查询
	_json_excel, _ := _se_excel.GetJson()
	fmt.Printf("当前excel---json是%s\n", _json_excel)

	//csv操作示例
	_se_csv := source.DocEntry{
		Type:      source.Csv,
		QueryType: source.GetList,
		Auth: source.DbAuth{
			DbName: "./test2.csv",
		},
	}
	//csv list查询
	_json_csv, _ := _se_csv.GetJson()
	fmt.Printf("当前csv---json是%s\n", _json_csv)
	//csv操作示例
	_se_txt := source.DocEntry{
		Type:      source.Txt,
		QueryType: source.GetList,
		Auth: source.DbAuth{
			DbName: "./test2.txt",
		},
	}
	//csv list查询
	_json_txt, _ := _se_txt.GetJson()
	fmt.Printf("当前txt---json是%s\n", _json_txt)
}

```
