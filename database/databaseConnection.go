package database

import (
	"context"
	"fmt"
	env "github.com/ntphuy2001/golang-jwt-project/dotenvsingleton"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func DBinstance() *mongo.Client {
	mongoDb := env.GetDotenvInstance().GetMongodbURL()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoDb))
	if err != nil {
		log.Fatal(err)
	}
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("GOLANG_AUTHENTICATION_WITH_JWT").Collection(collectionName)
	return collection
}
