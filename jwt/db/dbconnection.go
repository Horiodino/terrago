package jwt

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	Mongodb := os.Getenv(MongoDB_URL)

	client, err := mongo.NewClient(options.Client().ApplyURI(Mongodb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection established")

	return client
}

var MongoClient *mongo.Client = DBinstance()

// it will return the collection of database
func Opencollection(client *mongo.Client, collectionname string) *mongo.Collection {
	var DBCollection *mongo.Collection = client.Database("").Collection(collectionname)

	return DBCollection
}
