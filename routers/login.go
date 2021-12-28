package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MarianoChun/clonTwitter/bd"
	"github.com/MarianoChun/clonTwitter/jwt"
	"github.com/MarianoChun/clonTwitter/models"
)

/* Login realiza el login */
func Login(w http.ResponseWriter, r *http.Request) {
	// Lo que vamos a devolver es el Response Writer
	w.Header().Add("contenido-type", "application/json")
	var t models.Usuario
	// Lee el body y cargo los datos en t
	// Los datos son email y password
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		// Devolvemos el error al Response writer
		// Error 400 = BAD REQUEST
		http.Error(w, "Usuario y/o contraseña invalidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}
	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		http.Error(w, "Usuario y/o contraseña invalidos", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el Token correspondiente "+err.Error(), 400)
		return
	}

	respuesta := models.Respuestalogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// Codificamos a respuesta, que contenia un json en donde venia el Token
	json.NewEncoder(w).Encode(respuesta)

	// Aqui grabamos una cookie desde el back-end
	// Grabamos el jwt (json web token) en la cookie del usuario
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
