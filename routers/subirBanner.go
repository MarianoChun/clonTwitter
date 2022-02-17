package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/MarianoChun/clonTwitter/bd"
	"github.com/MarianoChun/clonTwitter/models"
)

/* SubirBanner sube el Banner al servidor */
func SubirBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	if err != nil {
		http.Error(w, "Error al obtener el archivo "+err.Error(), http.StatusBadRequest)
		return
	}
	// capturo la extension del archivo
	var extension = strings.Split(handler.Filename, ".")[1]
	// grabamos el avatar del usuario identificandolo con el ID de usuario (nunca se repite)
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension
	// Le doy permisos de lectura, escritura y ejecucion
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Banner = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || !status {
		http.Error(w, "Error al grabar el banner en la BD "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
