package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectMongo() {
	uri := "mongodb://root:root@localhost:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Erro conectando ao MongoDB:", err)
	}

	// Testa a conexão
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Erro ao pingar MongoDB:", err)
	}

	fmt.Println("✅ Conectado ao MongoDB")
	Client = client
}
