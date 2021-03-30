package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/AbduvokhidovRustamzhon/quote/pkg/model"
	"github.com/julienschmidt/httprouter"
)


func (server *server) handleCreateQuote(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	quote := model.Quote{}

	err := json.NewDecoder(request.Body).Decode(&quote)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	quote.CreatedAt = time.Now()
	err = server.quotes.CreateQuote(&quote)
	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	SendResponse(writer, quote)
}

func (server *server) handlerEditQuote(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	quote := &model.Quote{}
	err := json.NewDecoder(request.Body).Decode(&quote)
	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	editQuote, err := server.quotes.EditQuote(quote)
	if err != nil {
		log.Print(err)
		http.Error(writer, "id not exist", http.StatusNotFound)
		return
	}
	SendResponse(writer, editQuote)
}


func (server *server) handleRemoveQuote(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	quotes, err := server.quotes.Delete(id)
	if err == false {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	SendResponse(writer, quotes)
}