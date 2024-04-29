package main

import (
	"fmt"
	"os"
)

func Reading() {
	f, err := os.Open("sample.txt")
	defer f.Close()

	res := make([]byte, 1024)
	cnt, err := f.Read(res)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Read %d bytes!\n", cnt)
	fmt.Println(string(res[:cnt]))
}
