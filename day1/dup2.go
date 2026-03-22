package main

import (
	"bufio"
	"fmt"
	"os"
)

type Record struct {
	count    int
	filename string
}

func main() {
	counts := make(map[string]Record)
	// 读取命令后面的参数
	args := os.Args[1:]
	if len(args) == 0 {
		// 从标准输入读取
		countLines(os.Stdin, counts, "")
	} else {
		//从文件中读取
		for _, arg := range args {
			file, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(file, counts, arg)
			file.Close()
		}
	}
	for line, r := range counts {
		if r.count > 1 {
			if r.filename != "" {
				fmt.Println(line, "次数：", r.filename)
			}
		}
	}
}

func countLines(file *os.File, counts map[string]Record, filename string) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()] = Record{counts[input.Text()].count + 1, filename}
	}
}
