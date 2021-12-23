package bd

import "golang.org/x/crypto/bcrypt"

func EncriptarPassword(pass string) (string, error) {
	// El costo es un numero que determina la cantidad de pasadas sobre el texto para encriptarlo
	// El algoritmo utilizado esta basado en 2 ** costo para la cantidad de pasadas
	// A mayor costo, mayor cant de pasadas y mas demora.
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
