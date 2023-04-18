package utils

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

//var (
//	mgoCli        *mongo.Client
//	mgoCollection *mongo.Collection
//)
var client *mongo.Client

func InitEngine() *mongo.Client {
	cwdPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	err = godotenv.Load(cwdPath + "/.env")

	if err != nil {
		panic(err)
	}

	log.Println("Initializing connection to mongodb")

	log.Println(os.Getenv("MONGODB_CONNECTION_STRING"))

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING")).
		SetServerAPIOptions(serverAPIOptions)
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	log.Println("connected to mongodb")
	return client
}

//func GetMgoCli() *mongo.Client {
//	if mgoCli == nil {
//		InitEngine()
//	}
//	return mgoCli
//}

func GetCollection(databaseName, collectionName string) *mongo.Collection {
	if client == nil {
		client = InitEngine()
	}

	mgoCollection := client.Database(os.Getenv(databaseName)).Collection(os.Getenv(collectionName))
	return mgoCollection
}
