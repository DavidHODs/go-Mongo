package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/DavidHODs/go-Mongo/model"
	"github.com/DavidHODs/go-Mongo/router"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


const portNumber = ":8080"

var client *mongo.Client

func DummyEndPoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var person model.Person
	json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("myDb").Collection("people")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("\nStarting the application....")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	routes := router.Router()

	srv := &http.Server{
		Addr:  portNumber,
		Handler: routes,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		}
	
	err := srv.ListenAndServe()
	log.Fatal(err)
}

