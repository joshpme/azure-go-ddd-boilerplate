package main

import (
	"ddd-boilerplate/internal/controller"
	"ddd-boilerplate/internal/database"
	"github.com/akamensky/argparse"
	"log"
	"net/http"
	"os"
)

func main() {
	parser := argparse.NewParser("app", "Starts a webserver to serve EMS API")

	databaseConnection := parser.String("d", "database-connection-string", &argparse.Options{Required: true, Help: "Connection string to Cosmos DB"})
	err := parser.Parse(os.Args)
	if err != nil {
		log.Fatal(parser.Usage(err))
	}
	storageHandler := database.GetInstance(*databaseConnection)
	instrumentManager := controller.Connections{DbManager: storageHandler}

	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	http.HandleFunc("/api/instruments", instrumentManager.List)
	http.HandleFunc("/api/instruments/locations", instrumentManager.Locations)
	http.HandleFunc("/api/instruments/create", instrumentManager.Create)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
