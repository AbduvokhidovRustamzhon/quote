package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	_ "github.com/joho/godotenv/autoload"
	"github.com/AbduvokhidovRustamzhon/quote/cmd/app"
	"github.com/AbduvokhidovRustamzhon/quote/pkg/services"
	"github.com/julienschmidt/httprouter"
)


func main() {
	port, found := os.LookupEnv("SERVER_PORT")
	if !found {
		log.Fatal("Error, SERVER_PORT not set")
	}

	router := httprouter.New()
	quoteSvc := services.NewQuotes()
	server := app.NewServer(router, quoteSvc)
	server.Init()

	svc := http.Server{
		Handler: server,
		Addr:    port}

	fmt.Println("Server is listening on port", port)
	err := svc.ListenAndServe()
	if err != nil {
		log.Fatal("Error to run Server:", err)
	}
}