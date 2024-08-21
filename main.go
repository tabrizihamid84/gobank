package main

import (
	"fmt"
	"log"
)

func main() {

	store, err := NewPostgressStore()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewApiServer(":4000", store)
	server.Run()

	fmt.Println("Yeah Buddy!")
}
