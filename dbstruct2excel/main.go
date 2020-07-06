package main

import (
	"dbstruct2excel/config"
	"dbstruct2excel/dao"
	"dbstruct2excel/excel"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	mySQLConfig := config.GetConfig().MyMySQLConfig

	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mySQLConfig.Username, mySQLConfig.Password, mySQLConfig.Host, mySQLConfig.Port, mySQLConfig.DBName))
	if err != nil {
		log.Panicln(err)
	}

	saveByDDL(db, mySQLConfig.DBName)
	// saveByDML(db, mySQLConfig.DBName)
}

func saveByDML(db *sqlx.DB, dbName string) {
	tableInfos := dao.GetDMLTableInfo(db, dbName)
	for _, tableInfo := range tableInfos {
		header := &excel.Header{
			Name:    tableInfo.TableName,
			Comment: tableInfo.TableComment.String,
		}
		excel.WriteIndexSheet(header)

		filedInfos := dao.GetDMLFieldInfo(db, dbName, tableInfo.TableName)
		excelFieldInfos := []excel.FieldInfo{}
		for _, filedInfo := range filedInfos {
			excelFieldInfos = append(excelFieldInfos, excel.FieldInfo{
				Name:    filedInfo.FieldName,
				Type:    filedInfo.Type,
				Key:     filedInfo.Key.String,
				Null:    filedInfo.Null.String,
				Default: filedInfo.Default.String,
				Comment: filedInfo.Comment.String,
			})
		}

		excel.WriteTableInfo(header, excelFieldInfos)
	}

	excel.Save(dbName)
}

func saveByDDL(db *sqlx.DB, dbName string) {
	tableInfos := dao.GetDDLTableInfo(db)
	for _, tableInfo := range tableInfos {
		header := &excel.Header{
			Name:    tableInfo.Name,
			Comment: tableInfo.Comment.String,
		}
		excel.WriteIndexSheet(header)

		filedInfos := dao.GetDDLFieldInfo(db, tableInfo.Name)
		excelFieldInfos := []excel.FieldInfo{}
		for _, filedInfo := range filedInfos {
			excelFieldInfos = append(excelFieldInfos, excel.FieldInfo{
				Name:    filedInfo.Field,
				Type:    filedInfo.Type,
				Key:     filedInfo.Key.String,
				Null:    filedInfo.Null.String,
				Default: filedInfo.Default.String,
				Comment: filedInfo.Comment.String,
			})
		}

		excel.WriteTableInfo(header, excelFieldInfos)
	}

	excel.Save(dbName + "_" + time.Now().Format("2006-01-02"))
}
