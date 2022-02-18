package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/MarianoChun/clonTwitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* ConsultoRelacion consulta la relacion entre 2 usuarios */
func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	// Creo una condicion de busqueda (para encontrar la relacion)
	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}
	// Creamos un modelo de relacion
	var resultado models.Relacion
	fmt.Println(resultado)
	// Buscamos la relacion, pasando el contexto y la condicion de busqueda
	// y la pasamos a la variable "resultado"
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
