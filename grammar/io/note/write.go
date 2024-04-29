package main

import (
	"fmt"
	"os"
)

func Writing() {
	f, err := os.Create("sample.txt")

	str := "Hello, World!"
	res := []byte(str)
	cnt, err := f.Write(res)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Write %d bytes!\n", cnt)
}
