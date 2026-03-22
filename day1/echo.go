package main

import (
	"fmt"
	"os"
	"strconv"
)

func echo_with_index() {
	for i, s := range os.Args[1:] {
		fmt.Println("索引："+strconv.Itoa(i), "字符："+s)
	}
}
