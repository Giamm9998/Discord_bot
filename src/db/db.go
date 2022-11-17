package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Connect() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	_client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	client = _client
}

func Disconnect() {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	client = nil
}

func WriteCVE(cveData [][4]string) {
	coll := client.Database("News").Collection("CVEs")
	// We use update with option Upsert=true so the document is created only if it doesn't exist
	for _, cve := range cveData {
		update := bson.D{{Key: "$set", Value: bson.D{{Key: "id", Value: cve[0]}, {Key: "vuln", Value: cve[1]}, {Key: "score", Value: cve[2]}, {Key: "link", Value: cve[3]}}}}
		filter := bson.D{{Key: "id", Value: cve[0]}}
		opts := options.Update().SetUpsert(true)
		_, err := coll.UpdateOne(context.TODO(), filter, update, opts)
		if err != nil {
			panic(err)
		}
	}
}

//TODO: implement writeNews
