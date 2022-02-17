package bd

import (
	"context"
	"time"

	"github.com/MarianoChun/clonTwitter/models"
)

/* BorroRelacion borra la relacion en la base de datos*/
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")
	// Removemos la relacion
	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
