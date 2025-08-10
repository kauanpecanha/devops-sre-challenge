package db

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

var MongoClient *mongo.Client
var RollsCollection *mongo.Collection

func ConnectMongo() {
	godotenv.Load(".env")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Erro ao conectar no MongoDB: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Erro ao pingar MongoDB: %v", err)
	}

	fmt.Println("✅ Conectado ao MongoDB")
	MongoClient = client
	RollsCollection = client.Database("rolldice").Collection("rolls")
}
