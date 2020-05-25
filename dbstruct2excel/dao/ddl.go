package dao

import (
	"dbstruct2excel/models"
	"log"

	"github.com/jmoiron/sqlx"
)

// GetDDLTableInfo get ddl table info
func GetDDLTableInfo(db *sqlx.DB) []models.TableInfo {
	tableInfos := []models.TableInfo{}
	if err := db.Select(&tableInfos, "show table status"); err != nil {
		log.Panicln(err)
	}

	log.Printf("get %v table", len(tableInfos))

	return tableInfos
}

// GetDDLFieldInfo get ddl field infos
func GetDDLFieldInfo(db *sqlx.DB, tableName string) []models.FieldInfo {
	fieldInfos := []models.FieldInfo{}
	if err := db.Select(&fieldInfos, "show full columns from "+tableName); err != nil {
		log.Panicln(err)
	}

	log.Printf("%s table has %v field", tableName, len(fieldInfos))

	return fieldInfos
}
