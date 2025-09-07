package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// file's size
// `*` has higher priority than `<<`
const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
)

// generate file
func GenerateFile(path string, fileName string) *os.File {
	// 1. handle path
	// if path is not exist, create it with default folder permission
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0o0755)
		if err != nil {
			fmt.Println(err)
		}
	}

	// 2. create a new file
	name := filepath.Join(path, fileName)
	file, e := os.Create(name)
	if e != nil {
		fmt.Println(nil)
	}

	defer file.Close()

	return file
}

// write specific content and size to a file
func Writecontent(file *os.File, content []byte, size int) {

}

// calculate content's length
// according to the length, calculate how many times need to write in the specific size buffer
// output the exacltly lenght of buf
func Printcount(content string) {
	// Write content to a specific size buffer

}

func main() {
	content := `
		No matter what they tell us, no matter what they do, no matter what they teach us, what we believe is true.
		If I can see it, then I can do it. If I just believe it, there s nothing to it.
		In my dreams I always see you soar above the sky. In my heart there will always be a place for you and for my life. I keep a part of you with me and everywhere I am there you ll be
		I ll give you everything I am and everything I want to be. I put it in your hands if you could open up to me. Oh, can t we ever get beyong this wall. Because all I want is just once to see you in the light. But you hide behind the color of the night.
	`

	// test content
	// file1 := GenerateFile("./files", "1.txt")
	// file1.WriteString(content)

	fmt.Println(len(content))
}
