package models

// Tenemos 2 usuarios que se relacionan.
// 1ero Mi usuario y luego, el 2do usuario que voy a seguir

/* Relacion modelo para grabar la relacion de un usuario con otro*/
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioId"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionId"`
}
