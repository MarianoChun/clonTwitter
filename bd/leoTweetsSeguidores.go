package bd

import (
	"context"
	"time"

	"github.com/MarianoChun/clonTwitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	// Utilizo la coleccion de "relacion" ya que debo ver a quien sigo para visualizar sus tweets
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20

	// Creamos un slice de bson.M a la cual le iremos agregando condiciones
	condiciones := make([]bson.M, 0)
	// match busca el usuarioID de mi relacion
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	// El lookup me permite unir 2 tablas de la BD
	// Quiero que mis resultados obtenidos de "relacion" esten unidos a la tabla "tweets"
	condiciones = append(condiciones, bson.M{
		// Estoy uniendo mi tabla "relacion" a "tweet"
		// La voy a unir a traves del campo local "usuariorelacionid"
		// Voy a unir la tabla tweet en base a nuestros usuarios con los que estamos relacionados (userid)
		// Guardamos nuestra tabla con el mismo nombre "tweet"
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	// Con el "unwind" vamos a poder obtener los documentos de tal forma que sean todos iguales (con info repetida)
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	// Ordenamos los documentos de una forma determinada (fecha descendente)
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	// Limitamos la cantidad de tweets que visualizaremos (para no ver toda la bd de tweets)
	condiciones = append(condiciones, bson.M{"$limit": 20})
	// Con el Aggregate, vamos a poder procesar el "cursor" entero, sin recorrer cada uno de sus elementos
	cursor, _ := col.Aggregate(ctx, condiciones)
	var result []models.DevuelvoTweetSeguidores
	// Decodificamos todo el "cursor" en result
	err := cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}

	return result, true
}
