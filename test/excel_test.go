package test

import (
	"fmt"
	"github.com/silentryan/bigfiles/excel"
	"testing"
)

func TestReadAll(t *testing.T) {
	excel.ReadAll("C:/Users/L115840/Desktop/Excels")
}

// func TestCombine(t *testing.T) {
// 	paths := excel.ReadAll("C:/Users/L115840/Desktop/Excels")
// 	excel.Combinefile(paths)
// }

func TestHandleSingle(t *testing.T) {
	paths := excel.ReadAll("C:/Users/L115840/Desktop/Excels")
	fmt.Println(paths[0])
	rows := excel.HandleSingle(paths[0])
	fmt.Println(rows[0])
}

func TestWriteContent(t *testing.T) {
	excel.WriteExcel("C:/Users/L115840/Desktop/Excels", "total.xlsx")
}
