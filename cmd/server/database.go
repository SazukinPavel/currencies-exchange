package main

import (
	"context"
	"currencies-exchange/gen/currencies/v1"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDatabase struct {
	database_client *mongo.Client
	database        *mongo.Database
	cancel          context.CancelFunc
}

func new_mongo_database() *MongoDatabase {
	return &MongoDatabase{}
}

func (database *MongoDatabase) connect() {
	password := os.Getenv("MONGODB_ROOT_PASSWORD")
	username := os.Getenv("MONGO_ROOT_USERNAME")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	driver := os.Getenv("MONGO_DRIVER")
	database_name := os.Getenv("MONGO_DATABASE_NAME")

	var uri = driver + "://" + username + ":" + password + "@" + host + ":" + port + "/"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		panic(fmt.Sprintf("Mongo DB Connect issue %s", err))
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(fmt.Sprintf("Mongo DB ping issue %s", err))
	}

	database.database_client = client
	database.database = client.Database(database_name)
	database.cancel = cancel
}

func (database *MongoDatabase) close() {
	defer func() {
		database.cancel()
		if err := database.database_client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
		fmt.Println("Close connection is called")
	}()
}

type MongoCurrency struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
	Code string             `json:"code" bson:"code"`
}

type CurrencyCollection struct {
	collection *mongo.Collection
}

func new_currency_collection(database *MongoDatabase) *CurrencyCollection {
	return &CurrencyCollection{collection: database.database.Collection("currencies")}
}

// func (currency_collection CurrencyCollection) create(currency MongoCurrency) {

// 	currency.Id = primitive.NewObjectID()
// 	result, err := currency_collection.collection.InsertOne(context.TODO(), currency)
// 	if err != nil {
// 		fmt.Sprintf("Error in Creating currency %v", err)
// 		return
// 	}
// 	fmt.Printf("Inserte currency %v", result.InsertedID)
// }

// func (currency_collection CurrencyCollection) find_by_code(code string) MongoCurrency {

// 	filter := bson.D{{"code", code}}

// 	var result MongoCurrency
// 	currency_collection.collection.FindOne(context.TODO(), filter).Decode(&result)

// 	fmt.Printf("Find by code result %v", result)

// 	return result
// }

func (currency_collection CurrencyCollection) get_all() []MongoCurrency {
	var ctx = context.TODO()

	cur, err := currency_collection.collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var result []MongoCurrency

	for cur.Next(ctx) {
		var currency MongoCurrency
		err := cur.Decode(&currency)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, currency)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

// func (currency_collection CurrencyCollection) delete(currency MongoCurrency) {

// 	result, err := currency_collection.collection.DeleteOne(context.TODO(), currency)
// 	if err != nil {
// 		fmt.Sprintf("Error in Deleting currency %v", err)
// 		return
// 	}
// 	fmt.Printf("Delete currency %v", result)
// }

func get_currency_from_mongo_currency(ob1 *MongoCurrency) *currencies.Currency {
	ob2 := &currencies.Currency{Code: ob1.Code, Name: ob1.Name}
	return ob2
}

type MongoCurrencyRate struct {
	Id     primitive.ObjectID            `json:"id" bson:"_id"`
	From   string                        `json:"from" bson:"from"`
	To     string                        `json:"to" bson:"to"`
	Rate   float64                       `json:"rate" bson:"rate"`
	Status currencies.CurrencyRateStatus `json:"status" bson:"status"`
}

type RatesCollection struct {
	collection *mongo.Collection
}

func new_rates_collection(database *MongoDatabase) *RatesCollection {
	return &RatesCollection{collection: database.database.Collection("rates")}
}

func (rates_collection RatesCollection) get_all() []MongoCurrencyRate {
	var ctx = context.TODO()

	cur, err := rates_collection.collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var result []MongoCurrencyRate

	for cur.Next(ctx) {
		var rate MongoCurrencyRate
		err := cur.Decode(&rate)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, rate)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func get_rate_from_mongo_rate(ob1 *MongoCurrencyRate) *currencies.CurrencyRate {
	ob2 := &currencies.CurrencyRate{From: ob1.From, To: ob1.To, Rate: ob1.Rate, Status: ob1.Status}
	return ob2
}
