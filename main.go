package main

import (
	"fmt"
	"time"

	"fileHelp/file/cache"
)

func main() {
	cache.InitCache("cache")
	err := cache.Create("TEST", "Hello world!", time.Hour).Set()
	if err != nil {
		fmt.Println(err)
	}
	exist, value := cache.Get("TEST")
	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Error")
	}
	// cache.Flush("TEST")
}
