package main

import (
	"strings"
	"fmt"
	"io"
)

func testReader(reader io.Reader) ([]byte, error) {
	p := make([]byte, 6)
	n, err := reader.Read(p)
	if n > 0 {
		return p, nil
	}
	return nil, err
}
func testReaderAt(reader io.ReaderAt, offset int) []byte {
	if offset < 0 {
		return nil
	}
	p := make([]byte, offset)
	_, err := reader.ReadAt(p, 3)
	if err != nil {
		panic(err)
	}
	return p
}
func main() {
	reader := strings.NewReader("好好学习")
	bytes, err := testReader(reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("testReader %s\n", bytes)
	at := testReaderAt(reader, 6)

	fmt.Printf("testReaderAt %s", at)
}
