package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConnectBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://mUser:T0hD0yBG23w25wRb@cluster0.v1lr5.mongodb.net/test?retryWrites=true&w=majority")

func ConnectBD() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	clientOptions.SetMaxPoolSize(2000)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Conexión BD Ok")
	return client
}

func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err != nil
}
