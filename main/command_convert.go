package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/Soontao/go-project-template/lib"
	"github.com/urfave/cli"
)

var commandEntry = cli.Command{
	Name:   "convert",
	Usage:  "convert git-secrets output to excel",
	Action: convert,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "file, f",
			Usage:    "The git-secrets errors output file path",
			Required: true,
		},
		cli.StringFlag{
			Name:     "output, o",
			Usage:    "The output excel file",
			Required: true,
		},
	},
}

func convert(c *cli.Context) error {
	errorFilePath := c.String("file")
	output := c.String("output")
	if !strings.HasSuffix(output, ".xlsx") {
		output = fmt.Sprintf("%v.xlsx", output)
	}
	errorFile, err := os.OpenFile(errorFilePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer errorFile.Close()
	errorFileContent, err := ioutil.ReadAll(errorFile)
	if err != nil {
		return err
	}
	errorInfos := []*lib.LineErrorInformation{}
	for _, line := range strings.Split(string(errorFileContent), "\n") {
		errorInfo := lib.MatchLineString(line)
		if errorInfo != nil {
			errorInfos = append(errorInfos, errorInfo)
		}
	}

	if len(errorInfos) > 0 {
		f := excelize.NewFile()
		sheetName := "git-secrets-output"
		sheet := f.NewSheet(sheetName)
		f.SetActiveSheet(sheet)
		f.SetColWidth(sheetName, "A", "A", 50)
		f.SetColWidth(sheetName, "D", "D", 100)
		f.SetCellValue(sheetName, "A1", "File")
		f.SetCellValue(sheetName, "B1", "Line Number")
		f.SetCellValue(sheetName, "C1", "Status")
		f.SetCellValue(sheetName, "D1", "Matched Content")

		for idx, info := range errorInfos {
			f.SetCellValue(sheetName, fmt.Sprintf("A%v", idx+2), info.File)
			f.SetCellValue(sheetName, fmt.Sprintf("B%v", idx+2), info.Line)
			f.SetCellValue(sheetName, fmt.Sprintf("C%v", idx+2), "NEW")
			f.SetCellValue(sheetName, fmt.Sprintf("D%v", idx+2), info.Content)
		}

		if err != nil {
			return err
		}
		if err := f.SaveAs(output); err != nil {
			return err
		}
		log.Printf("save errors file to %v.\n", output)
	} else {
		log.Println("no errors found in git-secrets errors file.")
	}

	return nil

}
