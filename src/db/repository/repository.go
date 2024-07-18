package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// QuoteBD -
type QuoteBD struct {
	Name      string    `bson:"name"`
	Service   string    `bson:"service"`
	Deadline  int       `bson:"deadline"`
	Price     float64   `bson:"price"`
	CreatedAt time.Time `bson:"created_at"`
}

// CreateDB - creates the database
func CreateDB() {
	// create the database
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	// check if the connection is working
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	log.Println("Connected to MongoDB!")

	// create the collection
	quoteCollection := client.Database("freterapido").Collection("quotes")

	// insert a document
	var quote QuoteBD
	quote.CreatedAt = time.Now()

	result, err := quoteCollection.InsertOne(context.TODO(), quote)
	if err != nil {
		panic(err)
	}
	// display the id of the inserted entry
	fmt.Println(result.InsertedID)
}
