package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MarianoChun/clonTwitter/bd"
)

/* LeoTweetsSeguidores lee los tweets de todos nuestros seguidores*/
func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("pagina")
	if len(page) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina como entero mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := bd.LeoTweetsSeguidores(IDUsuario, pagina)
	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&respuesta)
}
