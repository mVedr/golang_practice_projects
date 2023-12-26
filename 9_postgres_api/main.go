package main

import (
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	err1 := store.Init()
	if err1 != nil {
		log.Fatal(err1)
	}
	//fmt.Printf("%+v \n", store)
	server := NewAPIServer(":3000", store)
	server.Run()
}
