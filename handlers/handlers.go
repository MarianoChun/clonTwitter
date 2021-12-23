package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/MarianoChun/clonTwitter/middlew"
	"github.com/MarianoChun/clonTwitter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Aqui se lista cada una de las rutas
/* Manejadores setea mi puerto, el handler y pongo a escuchar al servidor */
func Manejadores() {
	router := mux.NewRouter()
	// Se llama a registro, cuando detecta que se llama un POST, va a ejecutar el middleware
	// El middleware chequea la base de datos, si no hay problema de conexion, devuelve el control al router.

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
