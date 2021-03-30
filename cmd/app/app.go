package app

import (
	"encoding/json"
	"github.com/AbduvokhidovRustamzhon/quote/pkg/services/quote"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type server struct {
	router *httprouter.Router
	quotes *quote.Quotes
}

func NewServer(router *httprouter.Router, quotes *quote.Quotes) *server {
	return &server{router: router, quotes: quotes}
}

func (s *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
}

// SendResponse responses to client
func SendResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
