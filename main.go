package main

import (
	"context"
	"github.com/NisalSP9/Did-I-read/commons"
	"github.com/NisalSP9/Did-I-read/connections"
	"github.com/NisalSP9/Did-I-read/routes"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		commons.ErrorLogger.Println("Error loading .env file")
	}
	err = connections.CheckConnection(connections.Connect().Client()).Disconnect(context.TODO())
	if err != nil {
		commons.ErrorLogger.Println("Error disconnecting mongo connection")
	}
	headersOk := handlers.AllowedHeaders([]string{"Origin", "Range", "Content-Type", "Authorization",
		"Access-Control-Allow-Origin", "X-Accept-Language"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"DELETE", "GET", "HEAD", "POST", "PUT", "OPTIONS"})
	//Starting the API server
	router := routes.NewRouter()
	portNum := os.Getenv("DID_I_READ_ADMIN_BE_PORT")
	log.Println("Listening... port " + portNum)
	commons.ErrorLogger.Println(http.ListenAndServe(":"+portNum, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
