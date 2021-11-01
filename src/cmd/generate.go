/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"

	"github.com/gaku3601/auto-generate-sql/src/logic"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "指定ExcelからSQLを生成します",
	Long: `指定したExcelファイルからSQLを自動生成します。
-pで対象のExcelファイルを指定してください。`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		if err := logic.IsExistFile(path); err != nil {
			log.Fatal(err)
		}
		if err := logic.CheckExtension(path, []string{".xlsx", ".xlsm"}); err != nil {
			log.Fatal(err)
		}
		o, err := logic.NewOperationExcel(path)
		if err != nil {
			log.Fatal(err)
		}
		info := logic.ExtractDirPathAndName(path)
		if err := o.Execute(info.Path, info.Name); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("path", "p", "", "Excelのパスを指定")
	_ = generateCmd.MarkFlagRequired("path")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
