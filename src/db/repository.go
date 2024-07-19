package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"

	domain "frete-rapido/src/domain"
)

var mongoClient *mongo.Client

// QuoteBD -
type QuoteBD struct {
	Carriers  []Carrier `bson:"carrier"`
	CreatedAt time.Time `bson:"created_at"`
}

// Carrier -
type Carrier struct {
	Name     string  `bson:"name"`
	Service  string  `bson:"service"`
	Deadline int     `bson:"deadline"`
	Price    float64 `bson:"price"`
}

// CreateDB - creates the database
func CreateDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	// check if the connection is working
	if err := client.Ping(context.Background(), nil); err != nil {
		panic(err)
	}
	log.Println("Connected to MongoDB!")

	// set the client
	mongoClient = client
}

// SaveQuoteDB - saves the quote to the database
func SaveQuoteDB(responseQuote domain.ResponseQuote) {
	// create collection if it doesn't exist
	quoteCollection := mongoClient.Database("freterapido").Collection("quotes")

	// fill the object
	var quoteBD QuoteBD
	for _, carrier := range responseQuote.Carrier {
		quoteBD.Carriers = append(quoteBD.Carriers, Carrier{
			Name:     carrier.Name,
			Service:  carrier.Service,
			Deadline: carrier.Deadline,
			Price:    carrier.Price,
		})
	}
	quoteBD.CreatedAt = time.Now()

	// save the object
	quoteResult, err := quoteCollection.InsertOne(context.Background(), quoteBD)
	if err != nil {
		log.Printf("Error saving quote: %v", err)
		return
	}

	// get objectID
	log.Printf("Quote saved with ID: %v", quoteResult.InsertedID)

	return
}

// RetrieveQuotesDB - retrieves the quotes from the database
func RetrieveQuotesDB(lastQuotes int64) (quotes []QuoteBD, err error) {
	var cursor *mongo.Cursor

	// retrieve the collection
	quoteCollection := mongoClient.Database("freterapido").Collection("quotes")

	// check if lastquotes is set
	if lastQuotes > 0 {
		cursor, err = quoteCollection.Find(context.Background(), bson.M{}, options.Find().SetLimit(lastQuotes).SetSort(bson.M{"created_at": -1}))
	} else {
		cursor, err = quoteCollection.Find(context.Background(), bson.M{})
	}
	if err != nil {
		log.Printf("Error retrieving quotes: %v", err)
		return quotes, err
	}

	// convert the cursor to an array
	err = cursor.All(context.Background(), &quotes)
	if err != nil {
		log.Printf("Error converting quotes to json: %v", err)
		return quotes, err
	}

	return quotes, err
}
