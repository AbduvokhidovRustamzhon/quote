package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/AbduvokhidovRustamzhon/quote/pkg/model"
	"github.com/julienschmidt/httprouter"
)


func (server *server) handleCreateQuote(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	quote := model.Quote{}

	err := json.NewDecoder(r.Body).Decode(&quote)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	quote.CreatedAt = time.Now()
	err = server.quotes.CreateQuote(&quote)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	SendResponse(w, quote)
}