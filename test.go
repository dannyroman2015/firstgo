package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name   string
	Age    int
	Gender string `json:"gender`
}

func main() {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://dannyroman:ilove10111987@danny.augl8lb.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "skdfh ksdf")
	})

	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		collection := client.Database("danny").Collection("person")

		decoder := json.NewDecoder(r.Body)

		var p any
		err := decoder.Decode(&p)
		if err != nil {
			fmt.Println("here")
			panic(err)
		}
		fmt.Println(p)
		result, err := collection.InsertOne(context.TODO(), p)
		if err != nil {
			fmt.Println("o day")
			panic(err)
		}
		fmt.Println(result)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":8080", router)
}
