package bd

import (
	"github.com/MarianoChun/clonTwitter/models"
	"golang.org/x/crypto/bcrypt"
)

/* IntentoLogin realiza el chequeo de login a la BD*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usuario, encontrado, _ := ChequeoYaExisteUsuario(email)
	if !encontrado {
		return usuario, false
	}
	// Bcrypt trabaja con un slice de tipo byte

	// password no encriptada
	passwordBytes := []byte(password)
	// password encriptada
	passwordBD := []byte(usuario.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usuario, false
	}
	return usuario, true
}
