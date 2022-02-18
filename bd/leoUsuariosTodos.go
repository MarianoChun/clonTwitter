package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/MarianoChun/clonTwitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* LeoUsuariosTodos lee los usuarios registrados en el sistema, si se recibe "R" en quienes trae
solo los que se relacionan conmigo" */
func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var results []*models.Usuario
	// Seteamos nuestras opciones de busqueda (para visualizar la busqueda)
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	// Condicion de busqueda
	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}
	// Busco dentro de la coleccion de "usuarios" de acuerdo a las condiciones anteriormente definidas
	// la cual nos devuelve un "cursor" (Posee registros de usuarios)
	cursor, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var encontrado, incluir bool

	// Recorremos nuestro cursor
	// Por cada iteracion, veo si a ese usuario lo incluyo en una respuesta o no
	for cursor.Next(ctx) {
		var s models.Usuario
		err := cursor.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		// Consulto la relacion del usuario
		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		incluir = false

		encontrado, _ = ConsultoRelacion(r)

		// Si la condicion es cierta, incluyo al usuario en la lista
		// Si quiero listar a los que no sigo
		if tipo == "new" && !encontrado {
			incluir = true
		}
		// Si quiero listar a los que sigo
		if tipo == "follow" && encontrado {
			incluir = true
		}
		// Comprobamos que no nos estemos siguiendo a nosotros mismos para agregar a la lista
		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir {
			// Los campos que no quiero agregar a la lista, los paso aqui
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""
			// Agregamos el usuario a la lista
			results = append(results, &s)
		}
	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cursor.Close(ctx)
	return results, true
}
