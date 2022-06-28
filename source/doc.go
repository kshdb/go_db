package source

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xuri/excelize/v2"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
文档数据库操作对象
*/
type DocEntry struct {
	Type      int    `json:"type"`
	Auth      DbAuth `json:"auth"`
	QueryType int    `json:"query_type"`
}

/*
sql数据源
*/
func (s *DocEntry) GetJson() (_rt_json string, err error) {
	switch s.QueryType {
	case GetList:
		_rt_json, err = s.get_list()
		//fmt.Println("excel转json为",_rt_json)
		//case GetObject:
		//	_model,_err := s.get_one()
		//	err=_err
		//	_json, _err1 := json.Marshal(_model)
		//	err=_err1
		//	_rt_json = string(_json)
	}
	return
}

/*
列表查询
*/
func (s *DocEntry) get_list() (tableData string, err error) {
	switch s.Type {
	case Excel: //文档类excel数据库
		tableData = GetArryMap(s.Auth.DbName)
	case Txt, Csv: //文档类txt文本数据库
		data, _err := readAndParseCsv(s.Auth.DbName)
		if err != nil {
			err = _err
			fmt.Printf("处理csv文件时出错: %s\n", err)
		}
		tableData, err = csvToJson(data)
		if err != nil {
			fmt.Printf("将csv转换为json文件时出错: %s\n", err)
		}
		//fmt.Println(json)
	}
	return
}

/*
单条查询
*/
//func (s *DocEntry) get_one() (_one map[string]interface{},err error) {
//	//_one := map[string]interface{}{}
//	var _db Dbsql
//	switch s.Type {
//	case MySql:
//		_db = &DbMysql{Uid: s.Auth.Uid, Pwd: s.Auth.Pwd, IP: s.Auth.IP, Port: s.Auth.Port, DbName: s.Auth.DbName}
//	case MSSql:
//		_db = &DbMssql{Uid: s.Auth.Uid, Pwd: s.Auth.Pwd, IP: s.Auth.IP, Port: s.Auth.Port, DbName: s.Auth.DbName}
//	case Sqlite:
//		_db = &DbSqlite{Pwd: s.Auth.Pwd, DbName: s.Auth.DbName}
//	case Access:
//		_db = &DbAccess{Pwd: s.Auth.Pwd, DbName: s.Auth.DbName}
//	}
//	rows, _ := _db.Query(_sql)
//	defer rows.Close()
//	columns, _err1 := rows.Columns()
//	if _err1!=nil{
//		err=_err1
//		fmt.Println("数据读取出错了",_err1)
//	}
//	count := len(columns)
//	values := make([]interface{}, count)
//	scanArgs := make([]interface{}, count)
//	for i := range values {
//		scanArgs[i] = &values[i]
//	}
//	if rows.Next() {
//		rows.Scan(scanArgs...)
//		entry := make(map[string]interface{})
//		for i, col := range columns {
//			v := values[i]
//			b, ok := v.([]byte)
//			if ok {
//				entry[col] = string(b)
//			} else {
//				entry[col] = v
//			}
//		}
//		_one = entry
//	}
//	return
//}
//

/*-------------------------excel处理相关--------------------------------*/

type meta struct {
	Key string
	Idx int
	Typ string
}

type rowdata []interface{}

/*
获取json数组
*/
func GetArryMap(_file_path string) (_str_json string) {
	xlsx, err := excelize.OpenFile(_file_path)
	if err != nil {
		fmt.Println("打开excel文件出错", err.Error())
	}
	sheets := xlsx.GetSheetList()
	for _, s := range sheets {
		rows, err := xlsx.GetRows(s)
		if err != nil {
			return
		}
		if len(rows) < 5 {
			return
		}
		colNum := len(rows[1])
		//fmt.Println("col num:", colNum)
		metaList := make([]*meta, 0, colNum)
		dataList := make([]rowdata, 0, len(rows)-4)
		for line, row := range rows {
			switch line {
			//case 0: // sheet 名
			case 0: // col name
				for idx, colname := range row {
					//fmt.Println(idx, colname, len(metaList))
					metaList = append(metaList, &meta{Key: colname, Idx: idx})
				}
			//case 1: // data type
			//
			//	//fmt.Println("meta cot:%d, rol cot:%d", len(metaList), len(row))
			//	for idx, typ := range row {
			//		metaList[idx].Typ = typ
			//	}
			//case 2: // desc
			default: //>= 4 row data
				data := make(rowdata, colNum)

				for k := 0; k < colNum; k++ {
					if k < len(row) {
						data[k] = row[k]
					}
				}
				dataList = append(dataList, data)
			}
		}
		_str_json = toJson(dataList, metaList)
	}
	return
}

/*
二维数组转json
*/
func toJson(datarows []rowdata, metalist []*meta) string {
	ret := "["
	for _, row := range datarows {
		ret += "{"
		for idx, meta := range metalist {
			ret += fmt.Sprintf("\"%s\":", meta.Key)
			if meta.Typ == "string" {
				if row[idx] == nil {
					ret += "\"\""
				} else {
					ret += fmt.Sprintf("\"%s\"", gconv.String(row[idx]))
				}
			} else {
				if row[idx] == nil || row[idx] == "" {
					ret += "0"
				} else {
					ret += fmt.Sprintf("\"%s\"", gconv.String(row[idx]))
				}
			}
			ret += ","
		}
		ret = ret[:len(ret)-1]
		ret += "},"
	}
	ret = ret[:len(ret)-1]
	ret += "]"
	return ret
}

/*------------------------------txt,csv处理相关---------------------------------------*/
/*
读取csv文件
*/
func readAndParseCsv(path string) ([][]string, error) {
	csvFile, err := os.Open(path)
	if err != nil {
		//return nil, fmt.Errorf("error opening %s\n", path)
		fmt.Println("打开csv或txt文件出错", err.Error())
	}
	//判断当前编码
	//_bianma := determineEncodeing(csvFile)
	//fmt.Print("当前编码是", _bianma)

	var rows [][]string

	reader := csv.NewReader(csvFile)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return rows, fmt.Errorf("failed to parse csv: %s", err)
		}

		rows = append(rows, row)
	}

	return rows, nil
}

// 判断传输来的文本的字符集格式是什么
func determineEncodeing(r io.Reader) encoding.Encoding {
	peek, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		fmt.Println("出错了", err.Error())
	}
	determineEncoding, _, _ := charset.DetermineEncoding(peek, "")
	return determineEncoding
}

/*
csv转json
*/
func csvToJson(rows [][]string) (string, error) {
	var entries []map[string]interface{}
	attributes := rows[0]
	for _, row := range rows[1:] {
		entry := map[string]interface{}{}
		for i, value := range row {
			attribute := attributes[i]
			// split csv header key for nested objects
			objectSlice := strings.Split(attribute, ".")
			internal := entry
			for index, val := range objectSlice {
				// split csv header key for array objects
				key, arrayIndex := arrayContentMatch(val)
				if arrayIndex != -1 {
					if internal[key] == nil {
						internal[key] = []interface{}{}
					}
					internalArray := internal[key].([]interface{})
					if index == len(objectSlice)-1 {
						internalArray = append(internalArray, value)
						internal[key] = internalArray
						break
					}
					if arrayIndex >= len(internalArray) {
						internalArray = append(internalArray, map[string]interface{}{})
					}
					internal[key] = internalArray
					internal = internalArray[arrayIndex].(map[string]interface{})
				} else {
					if index == len(objectSlice)-1 {
						internal[key] = value
						break
					}
					if internal[key] == nil {
						internal[key] = map[string]interface{}{}
					}
					internal = internal[key].(map[string]interface{})
				}
			}
		}
		entries = append(entries, entry)
	}

	bytes, err := json.MarshalIndent(entries, "", "	")
	if err != nil {
		return "", fmt.Errorf("Marshal error %s\n", err)
	}

	return string(bytes), nil
}

func arrayContentMatch(str string) (string, int) {
	i := strings.Index(str, "[")
	if i >= 0 {
		j := strings.Index(str, "]")
		if j >= 0 {
			index, _ := strconv.Atoi(str[i+1 : j])
			return str[0:i], index
		}
	}
	return str, -1
}

/*-----------------------------------编码断言处理-----------------------------------------------------*/
type Charset string

const (
	UTF8    = Charset("UTF-8")
	UTF16   = Charset("UTF-16")
	GB18030 = Charset("GB18030")
	UNKNOWN = Charset("UNKNOWN")
)

//判断内容编码格式
func GetStrCoding(data []byte) Charset {
	if isUtf8(data) == true {
		return UTF8
	} else if isGBK(data) == true {
		return GB18030
	} else {
		return UNKNOWN
	}
}

//判断是否为中文编码
func isGBK(data []byte) bool {
	length := len(data)
	var i int = 0
	for i < length {
		if data[i] <= 0x7f {
			//编码0~127,只有一个字节的编码，兼容ASCII码
			i++
			continue
		} else {
			//大于127的使用双字节编码，落在gbk编码范围内的字符
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}

//判断是否为utf-8编码
func isUtf8(data []byte) bool {
	i := 0
	for i < len(data) {
		if (data[i] & 0x80) == 0x00 {
			// 0XXX_XXXX
			i++
			continue
		} else if num := preNUm(data[i]); num > 2 {
			// 110X_XXXX 10XX_XXXX
			// 1110_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_0XXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_10XX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_110X 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// preNUm() 返回首个字节的8个bits中首个0bit前面1bit的个数，该数量也是该字符所使用的字节数
			i++
			for j := 0; j < num-1; j++ {
				//判断后面的 num - 1 个字节是不是都是10开头
				if (data[i] & 0xc0) != 0x80 {
					return false
				}
				i++
			}
		} else {
			//其他情况说明不是utf-8
			return false
		}
	}
	return true
}
func preNUm(data byte) int {
	var mask byte = 0x80
	var num int = 0
	//8bit中首个0bit前有多少个1bits
	for i := 0; i < 8; i++ {
		if (data & mask) == mask {
			num++
			mask = mask >> 1
		} else {
			break
		}
	}
	return num
}
