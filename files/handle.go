package files

import (
	"fmt"
	"os"
)

// File size
const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
)

// Go []byte has its default vaule, the len(byte array) is specific
// When you generate specific size file of byte, you just need initialize a specific byte array
func GenerateByteFile(size int, output string) string {
	return "output"
}

// Generate string file
// content: Must limit the origin content is number and english character, because it is 1 byte
// size: 1MB; limit to 1MB
// out: the file path, must contains the file name
func GenerateM(content byte) []byte {
	storem := make([]byte, 0, MB)

	// WriteString
	// Windows's `\n` is CRLF == 2bytes
	for i := 0; i < MB; i = i + 4 {
		storem = append(storem, content)
		storem = append(storem, content)
		storem = append(storem, '\r')
		storem = append(storem, '\n')
	}

	return storem
}

// Create a 1GB file
func GenerateStringFileG(content byte, out string) {
	// create a file
	f, err := os.Create(out)
	if err != nil {
		fmt.Println(nil)
	}
	defer f.Close()

	for range 1024 {
		storem := GenerateM(content)
		f.WriteString(string(storem))
	}
}
