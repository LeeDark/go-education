package main

import (
	"log"
)

func main() {
	app := setEndpoints()

	log.Fatal(app.Listen(":3000"))
}
