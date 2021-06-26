package main

import (
	"fmt"

	"fileHelp/file/cache"
)

func main() {
	cache.InitCache("cache")
	exist, value := cache.Get("TEST")
	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Error")
	}
	err := cache.CreateForever("TEST", "Hello world!").Set()
	if err != nil {
		return
	}
}
