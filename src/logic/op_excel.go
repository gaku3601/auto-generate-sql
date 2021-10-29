package logic

import (
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

func (o OperationExcel) Execute() error {
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
