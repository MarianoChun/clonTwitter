package main

import (
	"log"

	"github.com/MarianoChun/clonTwitter/bd"
	"github.com/MarianoChun/clonTwitter/handlers"
)

func main() {
	// Si no hay conexion a la BD, arrojo error, sino, hago la conexion
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
