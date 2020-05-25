package dao

import (
	"dbstruct2excel/models"
	"log"

	"github.com/jmoiron/sqlx"
)

// GetDMLTableInfo get dml table info
func GetDMLTableInfo(db *sqlx.DB, dbName string) []models.SecTableInfo {
	tableInfos := []models.SecTableInfo{}
	if err := db.Select(&tableInfos,
		"SELECT TABLE_NAME, TABLE_COMMENT from information_schema.TABLES where TABLE_SCHEMA = '"+dbName+"' and TABLE_TYPE = 'BASE TABLE'",
	); err != nil {
		log.Panic(err)
	}

	log.Printf("get %v table", len(tableInfos))

	return tableInfos
}

// GetDMLFieldInfo get dml field infos
func GetDMLFieldInfo(db *sqlx.DB, dnName, tableName string) []models.SecFieldInfo {
	fieldInfos := []models.SecFieldInfo{}
	if err := db.Select(&fieldInfos, "SELECT COLUMN_NAME, IS_NULLABLE, COLUMN_TYPE, COLUMN_KEY, COLUMN_DEFAULT, COLUMN_COMMENT "+
		"from information_schema.`COLUMNS` "+
		"where TABLE_NAME = '"+tableName+"' and TABLE_SCHEMA = '"+dnName+"' order by ORDINAL_POSITION",
	); err != nil {
		log.Panic(err)
	}

	log.Printf("%s table has %v field", tableName, len(fieldInfos))

	return fieldInfos
}
