package main

import (
	"Mongodb/controllers"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := httprouter.New()
	client := getClient()
	uc := controllers.NewUserController(client)
	r.GET("/user/:id", uc.GetUser)
	http.ListenAndServe("localhost:8080", r)
}

func getClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Connected to MongoDB")
	return client
}
