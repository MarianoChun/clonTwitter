package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* BorroTweet borra un tweet determinado */
func BorroTweet(ID string, UserID string) error {
	// ID es el del tweet que voy a borrar, y el UserID es el del usuario
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	// Convertimos el ID del tweet a un objID
	objID, _ := primitive.ObjectIDFromHex(ID)

	// Creamos la condicion
	condicion := bson.M{
		"_id":    objID,
		"userid": UserID,
	}

	_, err := col.DeleteOne(ctx, condicion)
	return err
}
