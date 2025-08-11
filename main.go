package main

// boilerplate incomming from opentelemetry documentation at https://opentelemetry.io/docs/languages/go/getting-started/

import (
	"context"
	"kauanpecanha/devops-challenge/db"
	"kauanpecanha/devops-challenge/routes"
	"os"
	"os/signal"

	//"kauanpecanha/devops-challenge/configs/otel"
	"log"

	"github.com/joho/godotenv"

	"github.com/blackhorseya/golang-101/pkg/otelx"
)

func main() {
	// environment variables loading
	godotenv.Load(".env")

	_, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// mongodb connection
	db.ConnectMongo()

	// workflow name definition
	//name := os.Getenv("WORKFLOW_NAME")
	const name = "ROLLDICE"
	// otelcollector connection string definition
	//otelcollectorConn := os.Getenv("OTEL_COLLECTOR_CONN")
	const otelcollectorConn = "localhost:4317"

	// OpenTelemetry initializing
	err := otelx.SetupOTelSDK(context.Background(), otelcollectorConn, name)
	if err != nil {
		log.Printf("Failed to initialize OpenTelemetry: %v", err)
		return
	}
	defer func() {
		err = otelx.Shutdown(context.Background())
		if err != nil {
			log.Printf("Failed to shutdown OpenTelemetry: %v", err)
		}
	}()

	routes.NewHTTPHandler(name)

}
