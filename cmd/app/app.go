package app


import (
	"encoding/json"
	"net/http"

	"github.com/AbduvokhidovRustamzhon/quote/pkg/services"
	"github.com/julienschmidt/httprouter"
)

type server struct {
	router  *httprouter.Router
	quotes *services.Quotes
}

func NewServer(router *httprouter.Router, quotes *services.Quotes) *server {
   return &server{router: router, quotes: quotes}
}

func (server *server) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
   server.router.ServeHTTP(writer, request)
}

// Send Response to client
func SendResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}