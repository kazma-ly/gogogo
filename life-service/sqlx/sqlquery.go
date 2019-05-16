package sqlx

import (
	"database/sql"
	"errors"
	"life-service/reflectex"
	"reflect"
	"strings"
)

// Query 查询
func (bs *BaseSql) Query(any interface{}) error {
	rows, err := getRows(bs.sql, bs.args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	cs, err := rows.Columns()
	if err != nil {
		return err
	}

	if rows.Next() {
		err := makeResult(any, cs, rows)
		return err
	}
	return nil
}

// Querys 查询列表
func (bs *BaseSql) Querys(any interface{}) error {
	value := reflect.ValueOf(any)
	ttpe := value.Type()
	slice := reflectex.DeRef(ttpe)
	if slice.Kind() != reflect.Slice {
		return errors.New("need slice")
	}
	base := reflectex.DeRef(slice.Elem()) // 返回切片的内部类型

	rows, err := getRows(bs.sql, bs.args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	cs, err := rows.Columns()
	if err != nil {
		return err
	}

	nVal := reflect.Indirect(value)
	for rows.Next() {
		val := reflect.New(base).Interface()
		err := makeResult(val, cs, rows)
		if err != nil {
			return err
		}

		nVal.Set(reflect.Append(nVal, reflect.Indirect(reflect.ValueOf(val))))
	}
	return nil
}

// GetRows 获得rows
func getRows(sql string, args ...interface{}) (*sql.Rows, error) {
	CheckDB()
	stmt, err := mysqldb.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	return rows, err
}

// makeResult 从rows中获得一个结果
// dest 需要获得的结果
// cs 列的名字
// rows 结果集
func makeResult(dest interface{}, cs []string, rows *sql.Rows) error {
	rx := reflectex.New(dest)
	nf := rx.V.NumField()
	var vs = rx.GetDBEntity()

	// 从数据库中获取
	err := rows.Scan(vs...)
	if err != nil {
		return err
	}

	// 比较查询赋值复制
	for ind, c := range cs {
		val := vs[ind]
		for i := 0; i < nf; i++ {
			name := rx.FieldName(i)
			if !compareString(name, c) {
				continue
			}
			rx.SetValueWithAny(i, val)
		}
	}
	return nil
}

// compareString 比较结构体的变量名字和数据库的是不是一样
func compareString(entityStr, databaseStr string) bool {
	upEntityStr := strings.ToUpper(entityStr)
	upDatabaseStr := strings.ToUpper(databaseStr)
	if upEntityStr == upDatabaseStr { // 转换成全大些测试一次
		return true
	}

	noUnderscode := makeUnderscodeToUp(databaseStr)
	if entityStr == noUnderscode { // 下划线分割的形式
		return true
	}
	return false
}

// makeUnderscodeToUp 去除字符串中的"_" 并把后一个单词大写
func makeUnderscodeToUp(some string) string {
	res := ""
	needUp := false
	for _, r := range some {
		if r == '_' { // 跳过这次，下一次的字母改成大写
			needUp = true
		} else {
			if needUp == true {
				res += strings.ToUpper(string(r))
			} else {
				res += string(r)
			}

			needUp = false
		}
	}

	return strings.Title(res)
}
