package middlew

import (
	"net/http"

	"github.com/MarianoChun/clonTwitter/bd"
)

/* ChequeoBD es el middlew que me permite conocer el estado de la BD */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	// http.HandlerFunc es una funcion, por lo tanto se devuelve
	// una funcion
	return func(w http.ResponseWriter, r *http.Request) {
		// Si dio error, arrojo un http.Error y hago un return vacio
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
