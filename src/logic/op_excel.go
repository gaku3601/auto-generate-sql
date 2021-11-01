package logic

import (
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type OperationExcel struct {
	file   *excelize.File
	sheets []string
}

// NewOperationExcel constructor
func NewOperationExcel(path string) (*OperationExcel, error) {
	o := new(OperationExcel)
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}
	o.file = f
	for _, sheet := range o.file.GetSheetMap() {
		if sheet != "設定" {
			o.sheets = append(o.sheets, sheet)
		}
	}
	return o, nil
}

func (o OperationExcel) Execute(outputPath string, fileName string) error {
	// SQLファイルを作成する
	f, err := NewFile(fmt.Sprintf("%s/%s.%s", outputPath, fileName, "sql"))
	if err != nil {
		return err
	}
	defer f.Close()

	// シート毎に処理する
	for _, sheet := range o.sheets {
		rows := o.file.GetRows(sheet)
		var headers []string
		var values [][]string
		for i, row := range rows {
			if i == 0 {
				headers = analyzeHeaders(row)
			} else {
				values = append(values, analyzeValues(row, len(headers)))
			}
		}
		sqls := CreateInserts(sheet, headers, values)
		for _, sql := range sqls {
			if _, err := fmt.Fprintln(f.fp, sql); err != nil {
				return err
			}
		}
	}
	return nil
}

// 値が入っているもののみheaderとして認識して返却する
func analyzeHeaders(row []string) []string {
	var cols []string
	for _, col := range row {
		if len(col) > 0 {
			cols = append(cols, col)
		}
	}
	return cols
}

// headerの列の数を実際の値であると認識して返却する
func analyzeValues(row []string, headerCount int) []string {
	var cols []string
	for i, col := range row {
		if i < headerCount {
			cols = append(cols, deleteNewLineExcelCode(col))
		}
	}
	return cols
}

// 不要な改行コードを削除する
func deleteNewLineExcelCode(col string) string {
	return strings.Replace(col, "_x000D_", "", -1)
}
