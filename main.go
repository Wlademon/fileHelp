package main

import (
	"fmt"
	
	"fileHelp/file"
)

func main() {
	fileWrite := file.File{Name: "cache/init.set"}
	_, err := fileWrite.WriteNewLine("123")
	if err != nil {
		fmt.Println(err)
	}
	fileWrite.Clear()
}



