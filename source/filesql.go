package source

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/alexbrainman/odbc"
	"github.com/jmoiron/sqlx"
	_ "github.com/logoove/sqlite"
	//_"github.com/mutecomm/go-sqlcipher/v4"
)

/*
sqlite数据库对象
*/
type DbSqlite struct {
	DbName string
	Pwd    string
}

/*
sqlite数据源
*/
func (d *DbSqlite) db() (db *sqlx.DB, err error) {
	//key := "2DD29CA851E7B56E4697B0E1F08507293D761A05CE4D1B628663F411A8086D99"
	//dbname := fmt.Sprintf("test2.db?_pragma_key=x'%s'&_pragma_cipher_page_size=4096", "888888")
	//fmt.Println("这里是",dbname1)
	//dbname := fmt.Sprintf("%s?cache=shared&mode=memory", d.DbName)
	dbname := fmt.Sprintf("%s?_pragma_key=x'%s'&_pragma_cipher_use_hmac=off&_pragma_cipher_page_size=4096&_pragma_kdf_iter=256000", "./test.db", hex.EncodeToString([]byte("888888")))
	db, err = sqlx.Open("sqlite3", dbname)
	return
}
func (d *DbSqlite) Query(_sql string) (rows *sql.Rows, err error) {
	db, _err := d.db()
	err = _err
	defer db.Close()
	rows, err = db.Query(_sql)
	return
}

/*
Access数据库对象
*/
type DbAccess struct {
	DbName string
	Pwd    string
}

/*
access数据源
*/
func (d *DbAccess) db() (db *sql.DB, err error) {
	dbname := fmt.Sprintf(`driver={Microsoft Access Driver (*.mdb, *.accdb)};DBQ=%s;`, d.DbName)
	db, err = sql.Open("odbc", dbname)
	return
}
func (d *DbAccess) Query(_sql string) (rows *sql.Rows, err error) {
	db, _err := d.db()
	err = _err
	defer db.Close()
	rows, err = db.Query(_sql)
	return
}

/*
Excel数据库对象
*/
type DbExcel struct {
	DbName string
}

/*
excel数据源
*/
func (d *DbExcel) db() (db *sql.DB, err error) {
	dbname := fmt.Sprintf(`driver={Microsoft Access Driver (*.mdb, *.accdb)};DBQ=%s;`, d.DbName)
	db, err = sql.Open("odbc", dbname)
	return
}
func (d *DbExcel) Query(_sql string) (rows *sql.Rows, err error) {
	db, _err := d.db()
	err = _err
	defer db.Close()
	rows, err = db.Query(_sql)
	return
}
