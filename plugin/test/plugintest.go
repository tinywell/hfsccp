package main

import (
	"fmt"
	"os"
	"plugin"
)

func main() {
	_, err := plugin.Open("../databackscc.so.1.11.1")
	if err != nil {
		fmt.Println("error open plugin: ", err)
		os.Exit(-1)
	}
	fmt.Println("success")
}
