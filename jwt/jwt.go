package jwt

import (
	"time"

	"github.com/MarianoChun/clonTwitter/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/* GeneroJWT genera el encriptado con JWT */
func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("DesarrolloMariano_Golang")
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(24 * time.Hour).Unix(),
	}
	// El primer argumento es el header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
