package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

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

// handle single csv
func ReadSingleCsv(filename string) [][]string {
	// open file
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// CSV Reader
	csvReader := csv.NewReader(f)

	// using CSV Reader read file's rows
	rows, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println("reach here")
	// fmt.Println(rows[0])
	return rows
}
