package test

import (
	"testing"

	"github.com/silentryan/bigfiles/files"
)

// test generate a specific size of string file
func TestGenerateFile(t *testing.T) {
	files.GenerateStringFileG('b', "C:/Users/Admin/Desktop/files/02.txt")
}
