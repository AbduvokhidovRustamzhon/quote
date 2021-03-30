package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/AbduvokhidovRustamzhon/quote/cmd/app"
	"github.com/AbduvokhidovRustamzhon/quote/pkg/services/quote"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
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
	quoteSvc := quote.NewQuotes()
	server := app.NewServer(router, quoteSvc)
	server.Init()

	go quote.Worker(time.Minute*5, quoteSvc.DeleteOldQuotes(time.Hour))

	svc := http.Server{
		Handler: server,
		Addr:    port,
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		exitCh := make(chan os.Signal, 1)
		signal.Notify(exitCh, os.Interrupt, os.Kill)
		<-exitCh
		err := svc.Shutdown(context.Background())
		if err != nil {
			log.Println("server shutdown: ", err)
		}
		close(exitCh)
	}()

	fmt.Println("Server is start to listening on port:", port)
	err := svc.ListenAndServe()
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Error in running process server:", err)
		}
	}
	wg.Wait()
	fmt.Println("server gracefully shut downed")
}
