package middlew

import (
	"net/http"

	"github.com/MarianoChun/clonTwitter/routers"
)

/* ValidoJWT permite validar el JWT que nos viene en la peticion*/
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// La rutina va a recibir un token y nos va a devolver si el token es valido o no
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token ! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
