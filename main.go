package main

import (
	"log"

	"github.com/snetwork/bd"
	"github.com/snetwork/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion")
		return
	}
	handlers.Manejadores()
}
