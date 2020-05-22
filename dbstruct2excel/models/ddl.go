package models

import "database/sql"

// DDL query
type (
	// TableInfo table info
	TableInfo struct {
		Name          string         `db:"Name"`
		Engine        sql.NullString `db:"Engine"`
		Version       sql.NullInt64  `db:"Version"`
		RowFormat     sql.NullString `db:"Row_format"`
		Rows          sql.NullInt64  `db:"Rows"`
		AvgRowLength  sql.NullInt64  `db:"Avg_row_length"`
		DataLength    sql.NullInt64  `db:"Data_length"`
		MaxDataLength sql.NullInt64  `db:"Max_data_length"`
		IndexLength   sql.NullInt64  `db:"Index_length"`
		DataFree      sql.NullInt64  `db:"Data_free"`
		AutoIncrement sql.NullInt64  `db:"Auto_increment"`
		CreateTime    sql.NullTime   `db:"Create_time"`
		UpdateTime    sql.NullTime   `db:"Update_time"`
		CheckTime     sql.NullTime   `db:"Check_time"`
		Collation     sql.NullString `db:"Collation"`
		Checksum      sql.NullInt64  `db:"Checksum"`
		CreateOptions sql.NullString `db:"Create_options"`
		Comment       sql.NullString `db:"Comment"`
	}

	// FieldInfo field info
	FieldInfo struct {
		Field      string         `db:"Field"`
		Type       string         `db:"Type"`
		Collation  sql.NullString `db:"Collation"`
		Null       sql.NullString `db:"Null"`
		Key        sql.NullString `db:"Key"`
		Default    sql.NullString `db:"Default"`
		Extra      sql.NullString `db:"Extra"`
		Privileges sql.NullString `db:"Privileges"`
		Comment    sql.NullString `db:"Comment"`
	}
)
