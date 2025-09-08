package excel

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/silentryan/bigfiles/csv"
	"github.com/xuri/excelize/v2"
)

type Handler func(path string) [][]string

// return a file-path array
func ReadAll(path string) []string {
	// read folder
	files, err := os.ReadDir(path)
	if os.IsNotExist(err) {
		fmt.Println(err)
	}

	// combine path and put it into a []string
	apath := make([]string, 16)
	for _, name := range files {
		aname := filepath.Join(path, name.Name())
		// fmt.Println(aname)
		apath = append(apath, aname)
	}
	return apath
}

// read all excel files
// return a single file with all elements
func Combinefile(paths []string, handler Handler) string {
	// create a new excel file
	newFile := excelize.NewFile()
	sheetName := "Sheet1"
	defer newFile.Close()

	// Start writing data from row 1
	currentRow := 1

	// read all files and write into new file
	for _, path := range paths {
		rows := handler(path)

		// handle rows
		for _, row := range rows {
			for colIndex, cellValue := range row {
				cell, _ := excelize.CoordinatesToCellName(colIndex+1, currentRow)
				newFile.SetCellValue(sheetName, cell, cellValue)
			}
			currentRow++
		}
	}

	// Save the combined file
	outputPath := "combined.xlsx"
	if err := newFile.SaveAs(outputPath); err != nil {
		fmt.Printf("Error saving combined file: %v\n", err)
		return ""
	}

	return outputPath
}

// from a single absolute excel file path
// It cannot handle CSV files
// Only handle '.xlsx' & '.xlsm'
func HandleSingle(path string) [][]string {
	// 1. Read a single excel file
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", path, err)
		return nil
	}
	defer f.Close()

	// Get the first sheet name
	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		fmt.Printf("No sheets found in file %s\n", path)
	}
	sourceSheet := sheets[0]

	// read first sheet
	rows, err := f.GetRows(sourceSheet)
	if err != nil {
		fmt.Printf("Error reading rows from file %s: %v\n", path, err)
	}

	return rows
}

// Write content to a new file
// return filename
func WriteExcel(path string, filename string) string {
	// create newfile
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Create a new sheet.
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// Write content into new file
	paths := ReadAll(path)
	currentRow := 1
	for _, path := range paths {
		rows := csv.ReadSingleCsv(path)

		// write rows
		for _, row := range rows {
			for colIndex, cellValue := range row {
				cell, _ := excelize.CoordinatesToCellName(colIndex, currentRow)
				f.SetCellValue("Sheet1", cell, cellValue)
			}
			currentRow++
		}

	}

	// set active sheet of workbook
	f.SetActiveSheet(index)
	// save file
	output := filepath.Join(path, "all", filename)
	if err := f.SaveAs(output); err != nil {
		fmt.Println(err)
	}

	return output
}

// read specific column
func ReadSingleCol(colname string, filename string) []string {
	// open file
	// delay close file
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	colData := make([]string, 100000)
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
	}

	colNum, err := excelize.ColumnNameToNumber(colname)
	if err != nil {
		fmt.Println("Cannot read specific column")
	}

	// Read data from second rows
	for i := 1; i < len(rows); i++ {
		colData = append(colData, rows[i][colNum])
	}

	return colData
}
