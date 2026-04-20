package main

import (
	"fmt"
	"os"
)

// go mod init go-learn
// go mod tidy
func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
