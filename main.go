package main

import (
	"fmt"

	"fileHelp/file/cache"
)

func main() {

	cache.InitCache("cache")
	// err := cache.CreateForever("TEST", "GErttrrrfffdhjавываыаыаываыj").Set()
	// if err != nil {
	// 	return
	// }
	exist, value := cache.Get("TEST")
	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Error")
	}

	// fileWrite := file.File{Name: "cache/init.set"}
	// _, err := fileWrite.WriteNewLine("123")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fileWrite.Clear()
}
