package main

import "fmt"

func main() {
	server := NewApiServer(":4000")
	server.Run()

	fmt.Println("Yeah Buddy!")
}
