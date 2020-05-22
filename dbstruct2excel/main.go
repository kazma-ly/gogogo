package main

import (
	"dbstruct2excel/config"
	"dbstruct2excel/excel"
	"dbstruct2excel/tableinfo"
	"fmt"
	"log"

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
	tableInfos := tableinfo.GetDMLTableInfo(db, dbName)
	for _, tableInfo := range tableInfos {
		header := &excel.ExcelHeader{
			Name:    tableInfo.TableName,
			Comment: tableInfo.TableComment.String,
		}
		excel.WriteIndexSheet(header)

		filedInfos := tableinfo.GetDMLFieldInfo(db, dbName, tableInfo.TableName)
		excelFieldInfos := []excel.ExcelFieldInfo{}
		for _, filedInfo := range filedInfos {
			excelFieldInfos = append(excelFieldInfos, excel.ExcelFieldInfo{
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
	tableInfos := tableinfo.GetDDLTableInfo(db)
	for _, tableInfo := range tableInfos {
		header := &excel.ExcelHeader{
			Name:    tableInfo.Name,
			Comment: tableInfo.Comment.String,
		}
		excel.WriteIndexSheet(header)

		filedInfos := tableinfo.GetDDLFieldInfo(db, tableInfo.Name)
		excelFieldInfos := []excel.ExcelFieldInfo{}
		for _, filedInfo := range filedInfos {
			excelFieldInfos = append(excelFieldInfos, excel.ExcelFieldInfo{
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

	excel.Save(dbName)
}
