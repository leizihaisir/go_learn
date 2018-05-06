package main

import (
	"os"
	"bufio"
)

func main() {
	file, err := os.Open("/Users/zihailei/go/src/go_learn/file/writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(file)
	writer.Flush()
}
