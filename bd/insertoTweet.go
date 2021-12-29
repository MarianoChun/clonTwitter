package bd

import (
	"context"
	"time"

	"github.com/MarianoChun/clonTwitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertoTweet graba el Tweet en la BD */
func InsertoTweet(t models.GraboTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	registro := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}

	resultado, err := col.InsertOne(ctx, registro)
	if err != nil {
		return string(""), false, err
	}
	objID, _ := resultado.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
