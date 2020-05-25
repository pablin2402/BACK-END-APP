package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN objeto conectarbd*/
var MongoCN = conectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://pablo:PPR174064@cluster0-cx4ie.mongodb.net/test?retryWrites=true&w=majority")

/*FUNCION QUE PERMITE CONECTAR LA BD*/
func conectarBD() *mongo.Client {
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
	log.Println("CONEXION EXITOSA")
	return client
}

/*ChequeoConnection ping a la BD*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {

		return 0
	}
	return 1
}
