package db

import (
	"context"
	"fmt"
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

func Write_CVE(cveData [][4]string) {
	coll := client.Database("News").Collection("CVEs")
	cves := []interface{}{}
	for _, cve := range cveData {
		cves = append(cves, bson.D{{Key: "id", Value: cve[0]}, {Key: "vuln", Value: cve[1]}, {Key: "score", Value: cve[2]}, {Key: "link", Value: cve[3]}})
	}
	result, err := coll.InsertMany(context.TODO(), cves)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
