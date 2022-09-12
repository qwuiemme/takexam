package main

import (
	"log"

	"github.com/hnnngn/take-exam/internal/properties"
	"github.com/hnnngn/take-exam/internal/server"
)

func main() {
	srv := new(server.Server)

	err := srv.Run(properties.RunAddress)

	if err != nil {
		log.Fatal(err)
	}
}
