// generate excel related functions
package excel

import (
	_ "github.com/xuri/excelize/v2"
)

// The first row of a excel file
// it is special which shows the functionaility of each column
type headrow struct {
	content []any
}

type datarow struct {
	data []any
}

type Excel struct {
	head headrow
	data datarow
}

// Create an excel file
func NewExcel(path string, name string) {

}
