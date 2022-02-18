package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* DevuelvoTweetSeguidores es la estructura con la que devolveremos los tweets */
type DevuelvoTweetSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID         string             `bson:"usuarioid" json:"userID,omitempty"`
	UsuarioRelacionID string             `bson:"usuariorelacionid" json:"userRelationID,omitempty"`
	Tweet             struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
