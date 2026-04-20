package main

import (
	"fmt"
	"os"
)

func echo_for() {
	var sep, s string
	for _, s1 := range os.Args[1:] {
		s += sep + s1
		sep = " "
	}
	fmt.Println(s)
}
