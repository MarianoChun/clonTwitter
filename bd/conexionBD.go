package bd

// Recordar que para importar funciones a otros paquetes,
// estas func deben estar publicas
import (
	"context"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Obtener KEY
func viperEnvVariable(key string) string {
	// Configuramos el path del config file
	viper.SetConfigFile(".env")
	// Leemos el archivo configurado como config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error al leer el config file %s", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Error al convertir la key a string")
	}
	return value
}

var MongoCN = ConectarBD() // Aqui se ejecuta la conexion a la base de datos, devuelve la conexion en si misma
var URI string = "mongodb+srv://Mariano:" + viperEnvVariable("MONGO_TOKEN") + "@twitter.psjyi.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
var clientOptions = options.Client().ApplyURI(URI)

/* ConectarBD es la funcion que me permite conectar la base de datos */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa con la BD")
	return client
}

/* ChequeoConnection es el Ping a la BD */
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
