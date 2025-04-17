package main

import (
	"log"

	"github.com/Wefdzen/medods/api/router"
)

func main() {
	r := router.SetupRouter()
	log.Fatal(r.Run(":8080"))
}
