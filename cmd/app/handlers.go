package app

import (
	"encoding/json"
	"github.com/AbduvokhidovRustamzhon/quote/pkg/services/quote/models"
	"github.com/AbduvokhidovRustamzhon/quote/types"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func (s *server) createQuote(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var quoteRequest models.Request
	err := json.NewDecoder(request.Body).Decode(&quoteRequest)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	quote := models.NewQuote(quoteRequest.Author, quoteRequest.Quote, quoteRequest.Category)
	err = s.quotes.CreateQuote(quote)
	if err != nil {
		log.Printf("creating quote: %v", err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	SendResponse(writer, quote)
}

func (s *server) editQuote(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var quoteRequest models.Request
	err := json.NewDecoder(request.Body).Decode(&quoteRequest)
	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	id := params.ByName("id")
	if len(id) == 0 {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = s.quotes.EditQuote(id, &quoteRequest)
	if err != nil {
		log.Print(err)
		if err == types.ErrNotFound {
			http.Error(writer, http.StatusText(http.StatusGone), http.StatusGone)
			return
		}
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (s *server) deleteQuote(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	err := s.quotes.Delete(id)
	if err != nil {
		log.Print(err)
		if err == types.ErrNotFound {
			http.Error(writer, http.StatusText(http.StatusGone), http.StatusGone)
			return
		}
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (s *server) getAllQuotes(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	quotes, err := s.quotes.GetAllQuotes()
	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	SendResponse(writer, quotes)
}

func (s *server) getAllQuotesByCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	category := params.ByName("categorytype")

	quotes, err := s.quotes.GetAllQuotesByCategory(category)
	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	SendResponse(writer, quotes)
}

func (s *server) handleGetRandomQuote(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	quote, err := s.quotes.GetRandomQuote()
	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	SendResponse(writer, quote)
}
