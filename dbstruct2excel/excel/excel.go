package excel

import (
	"dbstruct2excel/models"
	"encoding/json"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type (
	// Header excel header
	Header struct {
		Name    string
		Comment string
	}

	// FieldInfo excel field list struct
	FieldInfo struct {
		Name    string
		Type    string
		Key     string
		Null    string
		Default string
		Comment string
	}
)

var (
	// excel file
	excelFile *excelize.File
	// excel sheet control
	indexSheetName = "index"
	tableSheetName = "tabel"
	indexSheetNum  = 0
	tableSheetNum  = 1
	// excel line control
	indexSheetIndex = 1
	tableSheetIndex = 1
)

func init() {
	excelFile = excelize.NewFile()

	// "index" sheet setting
	excelFile.SetSheetName(excelFile.GetSheetName(0), indexSheetName)
	excelFile.SetColWidth(indexSheetName, "A", "F", 20)

	// "table" sheet settng
	excelFile.SetActiveSheet(excelFile.NewSheet(tableSheetName))
	excelFile.SetColWidth(tableSheetName, "A", "A", 40)
	excelFile.SetColWidth(tableSheetName, "B", "E", 20)
	excelFile.SetColWidth(tableSheetName, "F", "F", 40)
}

// WriteIndexSheet write index sheet
func WriteIndexSheet(excelHeader *Header) {
	excelFile.SetActiveSheet(excelFile.GetSheetIndex(indexSheetName))

	currentIndexStr := strconv.Itoa(indexSheetIndex)
	writeHeader(indexSheetName, currentIndexStr, excelHeader)
	indexSheetIndex++
}

// WriteTableInfo write a table info
func WriteTableInfo(excelHeader *Header, fieldInfos []FieldInfo) {
	excelFile.SetActiveSheet(excelFile.GetSheetIndex(tableSheetName))

	// write header
	tableSheetIndexStr := strconv.Itoa(tableSheetIndex)
	writeHeader(tableSheetName, tableSheetIndexStr, excelHeader)
	// setting style
	excelFile.SetCellStyle(tableSheetName, "A"+tableSheetIndexStr, "F"+tableSheetIndexStr, getHeaderStyleString(excelFile))
	// setting link
	excelFile.SetCellHyperLink(indexSheetName, "A"+strconv.Itoa(indexSheetIndex-1), tableSheetName+"!"+"A"+tableSheetIndexStr+":F"+tableSheetIndexStr, "Location")
	tableSheetIndex++

	// write section
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
	for _, filedInfo := range fieldInfos {
		tableSheetIndex = writeOneLineExcel(excelFile, tableSheetName, tableSheetIndex, &filedInfo)
	}

	// set boder style
	vcell := "F" + strconv.Itoa(tableSheetIndex-1)
	excelFile.SetCellStyle(tableSheetName, hcell, vcell, getBorderStyleString(excelFile))
	tableSheetIndex++
}

// Save save excel file
func Save(fileName string) {
	tableSheetIndex = 1
	excelFile.SetActiveSheet(excelFile.GetSheetIndex(indexSheetName))
	excelFile.SaveAs(fileName + ".xlsx")
}

func writeHeader(sheetName, currentIndexStr string, val *Header) {
	excelFile.MergeCell(sheetName, "A"+currentIndexStr, "C"+currentIndexStr)
	excelFile.SetCellStr(sheetName, "A"+currentIndexStr, val.Name)
	excelFile.MergeCell(sheetName, "D"+currentIndexStr, "F"+currentIndexStr)
	excelFile.SetCellStr(sheetName, "D"+currentIndexStr, val.Comment)
}

func writeOneLineExcel(excelFile *excelize.File, sheetName string, postion int, filedInfo *FieldInfo) int {
	currentPostion := strconv.Itoa(postion)

	excelFile.SetCellStr(sheetName, "A"+currentPostion, filedInfo.Name)
	excelFile.SetCellStr(sheetName, "B"+currentPostion, filedInfo.Type)
	excelFile.SetCellStr(sheetName, "C"+currentPostion, filedInfo.Key)
	excelFile.SetCellStr(sheetName, "D"+currentPostion, filedInfo.Null)
	excelFile.SetCellStr(sheetName, "E"+currentPostion, filedInfo.Default)
	excelFile.SetCellStr(sheetName, "F"+currentPostion, filedInfo.Comment)

	return postion + 1
}

// boder style
func getBorderStyleString(excelFile *excelize.File) int {
	bs := models.Style{
		Border: []models.Border{
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
	style := models.Style{
		Border: []models.Border{
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
		Fill: models.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"009999"},
		},
		Font: &models.Font{
			Bold:  true,
			Color: "FFFFFF",
		},
	}

	styleByte, _ := json.Marshal(style)
	styleVal, _ := excelFile.NewStyle(string(styleByte))
	return styleVal
}
