package models

import "database/sql"

// DML query
type (
	// SecTableInfo table info
	SecTableInfo struct {
		TableName    string         `db:"TABLE_NAME"`
		TableComment sql.NullString `db:"TABLE_COMMENT"`
	}

	// SecFieldInfo field info
	SecFieldInfo struct {
		FieldName string         `db:"COLUMN_NAME"`
		Type      string         `db:"COLUMN_TYPE"`
		Key       sql.NullString `db:"COLUMN_KEY"`
		Null      sql.NullString `db:"IS_NULLABLE"`
		Default   sql.NullString `db:"COLUMN_DEFAULT"`
		Comment   sql.NullString `db:"COLUMN_COMMENT"`
	}
)
