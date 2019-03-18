package main

import (
	"os"
	"log"
	"bufio"
	"io"
	"fmt"
)

func main() {
	file, err := os.Open("/Users/zihailei/Documents/server.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		//val, err := reader.ReadString('\n') // 或者 reader.ReadLine
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))

	}
}
