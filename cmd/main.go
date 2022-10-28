package main

import (
	"log"
	"zuri-stage-one/cmd/server"
)

func main() {
	db, err := server.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
	server.Injection(db)
}
