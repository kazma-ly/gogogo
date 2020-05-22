package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gopkg.in/yaml.v2"
)

type (
	// Config project config
	Config struct {
		MyMySQLConfig MySQLConfig `yaml:"mysql"`
	}

	// MySQLConfig mysql config
	MySQLConfig struct {
		Host     string `yaml:"host"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Port     string `yaml:"port"`
		DBName   string `yaml:"dbname"`
	}

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

	// FiledInfo field info
	FiledInfo struct {
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

	// Font font style
	Font struct {
		Bold      bool    `json:"bold"`
		Italic    bool    `json:"italic"`
		Underline string  `json:"underline"`
		Family    string  `json:"family"`
		Size      float64 `json:"size"`
		Strike    bool    `json:"strike"`
		Color     string  `json:"color"`
	}

	// Fill fill style
	Fill struct {
		Type    string   `json:"type"`
		Pattern int      `json:"pattern"`
		Color   []string `json:"color"`
		Shading int      `json:"shading"`
	}

	// Border border style
	Border struct {
		Type  string `json:"type"`
		Color string `json:"color"`
		Style int    `json:"style"`
	}

	// Style style
	Style struct {
		Border []Border `json:"border"`
		Fill   Fill     `json:"fill"`
		Font   *Font    `json:"font"`
	}
)

func main() {

	configFile, err := os.OpenFile("./config.yml", os.O_RDONLY, 0755)
	if err != nil {
		log.Panicln(err)
	}

	configByte, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Panicln(err)
	}

	config := &Config{}
	if yaml.Unmarshal(configByte, config) != nil {
		log.Panicln(err)
	}

	mySQLConfig := config.MyMySQLConfig

	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mySQLConfig.Username, mySQLConfig.Password, mySQLConfig.Host, mySQLConfig.Port, mySQLConfig.DBName))

	if err != nil {
		log.Panicln(err)
	}

	tableInfos := []TableInfo{}
	if err := db.Select(&tableInfos, "show table status"); err != nil {
		log.Panicln(err)
	}

	log.Printf("get %v table", len(tableInfos))

	excelFile := excelize.NewFile()
	indexSheetName := "index"
	excelFile.SetSheetName(excelFile.GetSheetName(1), indexSheetName)
	excelFile.SetColWidth(indexSheetName, "A", "F", 20)

	// write index
	var indexSheetIndex = 1
	for _, tableInfo := range tableInfos {
		currentIndexStr := strconv.Itoa(indexSheetIndex)
		excelFile.MergeCell(indexSheetName, "A"+currentIndexStr, "C"+currentIndexStr)
		excelFile.SetCellStr(indexSheetName, "A"+currentIndexStr, tableInfo.Name)
		excelFile.MergeCell(indexSheetName, "D"+currentIndexStr, "F"+currentIndexStr)
		excelFile.SetCellStr(indexSheetName, "D"+currentIndexStr, tableInfo.Comment.String)
		indexSheetIndex++
	}

	// write table
	tableSheetName := "tabel"
	excelFile.SetActiveSheet(excelFile.NewSheet(tableSheetName))
	excelFile.SetColWidth(tableSheetName, "A", "A", 40)
	excelFile.SetColWidth(tableSheetName, "B", "E", 20)
	excelFile.SetColWidth(tableSheetName, "F", "F", 40)
	var tableSheetIndex = 1
	for i, tableInfo := range tableInfos {
		log.Printf("current: %v", tableSheetIndex)
		// write section
		tableSheetIndexStr := strconv.Itoa(tableSheetIndex)
		excelFile.MergeCell(tableSheetName, "A"+tableSheetIndexStr, "C"+tableSheetIndexStr)
		excelFile.SetCellStr(tableSheetName, "A"+tableSheetIndexStr, tableInfo.Name)
		excelFile.MergeCell(tableSheetName, "D"+tableSheetIndexStr, "F"+tableSheetIndexStr)
		excelFile.SetCellStr(tableSheetName, "D"+tableSheetIndexStr, tableInfo.Comment.String)
		// setting style
		excelFile.SetCellStyle(tableSheetName, "A"+tableSheetIndexStr, "F"+tableSheetIndexStr, getHeaderStyleString(excelFile))
		// setting link
		excelFile.SetCellHyperLink(indexSheetName, "A"+strconv.Itoa(i+1), tableSheetName+"!"+"A"+tableSheetIndexStr+":F"+tableSheetIndexStr, "Location")
		tableSheetIndex++

		// write header
		tableSheetIndexStr = strconv.Itoa(tableSheetIndex)
		hcell := "A" + tableSheetIndexStr
		excelFile.SetCellStr(tableSheetName, "A"+tableSheetIndexStr, "字段名称")
		excelFile.SetCellStr(tableSheetName, "B"+tableSheetIndexStr, "字段类型")
		excelFile.SetCellStr(tableSheetName, "C"+tableSheetIndexStr, "键")
		excelFile.SetCellStr(tableSheetName, "D"+tableSheetIndexStr, "是否允许为空")
		excelFile.SetCellStr(tableSheetName, "E"+tableSheetIndexStr, "默认值")
		excelFile.SetCellStr(tableSheetName, "F"+tableSheetIndexStr, "注释")
		tableSheetIndex++

		// write body
		filedInfos := []FiledInfo{}
		if err := db.Select(&filedInfos, "show full columns from "+tableInfo.Name); err != nil {
			log.Panicln(err)
		}

		log.Printf("%s table has %v field", tableInfo.Name, len(filedInfos))

		for _, filedInfo := range filedInfos {
			tableSheetIndex = writeOneLineExcel(excelFile, tableSheetName, tableSheetIndex, &filedInfo)
		}

		// set boder style
		vcell := "F" + strconv.Itoa(tableSheetIndex-1)
		excelFile.SetCellStyle(tableSheetName, hcell, vcell, getBorderStyleString(excelFile))

		tableSheetIndex++
	}

	excelFile.SetActiveSheet(1)

	excelFile.SaveAs(mySQLConfig.DBName + ".xlsx")
}

// writeOneLineExcel write one line
func writeOneLineExcel(excelFile *excelize.File, sheetName string, postion int, filedInfo *FiledInfo) int {
	currentPostion := strconv.Itoa(postion)

	excelFile.SetCellStr(sheetName, "A"+currentPostion, filedInfo.Field)
	excelFile.SetCellStr(sheetName, "B"+currentPostion, filedInfo.Type)
	excelFile.SetCellStr(sheetName, "C"+currentPostion, filedInfo.Key.String)
	excelFile.SetCellStr(sheetName, "D"+currentPostion, filedInfo.Null.String)
	excelFile.SetCellStr(sheetName, "E"+currentPostion, filedInfo.Default.String)
	excelFile.SetCellStr(sheetName, "F"+currentPostion, filedInfo.Comment.String)

	return postion + 1
}

// boder style
func getBorderStyleString(excelFile *excelize.File) int {
	bs := Style{
		Border: []Border{
			{
				Type:  "left",
				Color: "000000",
				Style: 1,
			}, {
				Type:  "top",
				Color: "000000",
				Style: 1,
			}, {
				Type:  "bottom",
				Color: "000000",
				Style: 1,
			}, {
				Type:  "right",
				Color: "000000",
				Style: 1,
			},
		},
	}

	bsByte, _ := json.Marshal(bs)
	borderStyle, _ := excelFile.NewStyle(string(bsByte))
	return borderStyle
}

// header style
func getHeaderStyleString(excelFile *excelize.File) int {
	style := Style{
		Border: []Border{
			{
				Type:  "left",
				Color: "000000",
				Style: 1,
			}, {
				Type:  "top",
				Color: "000000",
				Style: 1,
			}, {
				Type:  "bottom",
				Color: "000000",
				Style: 1,
			}, {
				Type:  "right",
				Color: "000000",
				Style: 1,
			},
		},
		Fill: Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"009999"},
		},
		Font: &Font{
			Bold:  true,
			Color: "FFFFFF",
		},
	}

	styleByte, _ := json.Marshal(style)
	styleVal, _ := excelFile.NewStyle(string(styleByte))
	return styleVal
}
