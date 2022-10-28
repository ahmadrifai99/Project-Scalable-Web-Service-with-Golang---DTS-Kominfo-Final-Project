package main

import (
	"final-project/infrastructure/server"
	"log"
)

func main() {
	app := server.InitServer()

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
