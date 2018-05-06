package main

import "os"

func main() {
	file, err := os.Create("file/writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var data = []byte("我爱go语言")
	file.Write(data)

	_, writeAterr := file.WriteAt([]byte("Go语言中文网"), 14)
	if writeAterr != nil {
		panic(err)
	}
}
