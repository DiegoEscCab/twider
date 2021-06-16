package main

import (
	"log"

	"github.com/DiegoEscCab/twider/tree/main/bd"
	"github.com/DiegoEscCab/twider/tree/main/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la base de datos")
		return
	}
	handlers.Manejadores()
}
