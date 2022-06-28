package source

import "fmt"

/*
nosql数据库类型
*/
const (
	Redis = iota + 201
	Mongo
)

/*
nosql数据源
*/
func DbNoSql(_type int) {
	switch _type {
	case Redis:
		fmt.Println("当前数据库是Redis", _type)
	case Mongo:
		fmt.Println("当前数据库是Mongo", _type)
	}
}
