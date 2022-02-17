package routers

import (
	"net/http"

	"github.com/MarianoChun/clonTwitter/bd"
	"github.com/MarianoChun/clonTwitter/models"
)

// BajaRelacion realiza el borrado de la relacion entre usuarios
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	// Obtenemos el ID del usuario que me quiero relacionar
	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID
	// Borramos la relacion
	status, err := bd.BorroRelacion(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado borrar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)

}
