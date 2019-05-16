package sqlx

import (
	"database/sql"
	"life-service/confs"
	"life-service/logx"
	"log"
	"sync"
)

var (
	mysqldb   *sql.DB
	mysqlLock sync.Mutex
)

// CheckDB 做数据库连接
func CheckDB() {
	mysqlLock.Lock()
	defer mysqlLock.Unlock()

	if mysqldb == nil {
		_mysqldb, err := sql.Open("mysql", confs.ReadConfValue("mysql.url").(string))
		if err != nil || _mysqldb == nil {
			log.Panicln("数据库连接失败", err)
		} else {
			mysqldb = _mysqldb
			mysqldb.SetMaxIdleConns(10)
			log.Println("数据库连接成功: ", mysqldb.Stats())
		}
	} else {
		if mysqldb.Ping() != nil {

		}
	}
}

func initDB() {
	_mysqldb, err := sql.Open("mysql", confs.ReadConfValue("mysql.url").(string))
	if err != nil {
		logx.LogInfo("数据库连接失败: " + err.Error())
	} else {
		mysqldb = _mysqldb
		mysqldb.SetMaxIdleConns(10)
		logx.LogInfo("数据库连接成功: ", mysqldb.Stats())
	}
}
