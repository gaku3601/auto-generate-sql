package logic

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type OperationExcel struct {
	file   *excelize.File
	sheets []string
	envs   map[string]string
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
		if sheet != "env" {
			o.sheets = append(o.sheets, sheet)
		}
	}
	o.envs = make(map[string]string)
	o.analyzeEnvs()
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
				headers = row
			} else {
				values = append(values, row)
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

// envシートが存在する場合、envの内容を格納する
func (o OperationExcel) analyzeEnvs() {
	rows := o.file.GetRows("env")
	for _, row := range rows {
		if len(row) != 2 {
			log.Fatal("環境変数はA列にkey、B列にvalueを必ず入力してください")
		}
		o.envs[row[0]] = row[1]
	}
}
