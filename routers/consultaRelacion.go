package routers

import (
	"encoding/json"
	"net/http"

	"github.com/MarianoChun/clonTwitter/bd"
	"github.com/MarianoChun/clonTwitter/models"
)

/* ConsultaRelacion chequea si hay relacion entre 2 usuarios */
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	// Creamos un modelo de relacion y completamos sus campos
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion

	status, err := bd.ConsultoRelacion(t)
	// Asignamos a "resp" true o false dependiendo de si hay relacion o no
	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}
