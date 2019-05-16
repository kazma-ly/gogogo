package sqlx

import (
	"database/sql"
)

type (
	// BaseSql 基础查询接口
	BaseSql struct {
		sql  string
		args []interface{}
	}
)

// New 初始化一个sql操作类
func New(sql string, args []interface{}) BaseSql {
	return BaseSql{sql: sql, args: args}
}

// Exec 执行命令
func (bs *BaseSql) Exec() (sql.Result, error) {
	CheckDB()
	tx, err := mysqldb.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	stmt, err := mysqldb.Prepare(bs.sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Exec(bs.args...)
}
