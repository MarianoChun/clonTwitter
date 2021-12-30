package bd

import (
	"context"
	"log"
	"time"

	"github.com/MarianoChun/clonTwitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* LeoTweets lee los tweets de un perfil */
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweets
	// Condicion a buscar en la BD
	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	// Esto va a sortear todos los documentos de acuerdo a la fecha en orden descendente
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	// De acuerdo al valor de pagina, voy salteando
	opciones.SetSkip((pagina - 1) * 20)
	// Cursor es un puntero (especie de tabla de BD), en el se graban los resultados y los puedo recorrer de a uno
	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true
}
