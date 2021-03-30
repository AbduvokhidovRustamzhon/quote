package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/AbduvokhidovRustamzhon/quote/cmd/app"
	"github.com/AbduvokhidovRustamzhon/quote/pkg/services"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

// init is invoked before main()
func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func main() {
	port, found := os.LookupEnv("SERVER_PORT")
	if !found {
		log.Fatal("Error, SERVER_PORT not set")
	}

	router := httprouter.New()
	quoteSvc := services.NewQuotes()
	server := app.NewServer(router, quoteSvc)
	server.Init()
	go services.Worker(time.Minute * 5, quoteSvc.DeleteOldQuotes)

	svc := http.Server{
		Handler: server,
		Addr:    port}

	fmt.Println("Servrer is start to listening on port:", port)
	err := svc.ListenAndServe()
	if err != nil {
		log.Fatal("Error in running process server:", err)
	}

	time.Sleep(time.Minute * 1)
}