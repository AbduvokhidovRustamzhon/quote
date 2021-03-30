package services

import (
	"sync"
	"github.com/AbduvokhidovRustamzhon/quote/constants"
	"github.com/AbduvokhidovRustamzhon/quote/pkg/model"
	"github.com/google/uuid"
)


type Quotes struct {
	Quotes map[string]model.Quote
	sync.Mutex
}

func NewQuotes() *Quotes {
	return &Quotes{Quotes: make(map[string]model.Quote)}
}

//Create Quotes
func (quotes *Quotes) CreateQuote(quote *model.Quote) (err error) {
	quotes.Lock()
	defer  quotes.Unlock()
	quote.ID = uuid.New().String()
	quotes.Quotes[quote.ID] = *quote

	if quotes.Quotes == nil {
		return err
	}
	return nil
}

// Edit Quote 
func (quotes *Quotes) EditQuote(quote *model.Quote) (*model.Quote, error) {
	quotes.Lock()
	defer  quotes.Unlock()
	for key, _ := range quotes.Quotes {
		if key == quote.ID {
			quotes.Quotes[quote.ID] = *quote
			return quote, nil
		}

	}
	return nil, constants.ErrIDNotFound
}


// Delete Quote
func (quote *Quotes) Delete(quoteID string) ([]model.Quote, bool) {

	_, exists := quote.Quotes[quoteID]
	if exists {
		delete(quote.Quotes, quoteID)
		quotes, _ := quote.GetAllQuotes()
		return quotes, true
	}
	return nil, false
}

// Get Quotes
func (quote *Quotes) GetAllQuotes() ([]model.Quote, error) {
	quote.Lock()
	defer  quote.Unlock()
	quotes := []model.Quote{}

	for _, value := range quote.Quotes {
		quotes = append(quotes, value)

	}
	if quotes == nil {
		return nil, constants.ErrNotFound
	}
	return quotes, nil
}

// Get All Quotes by Categore
func (quote *Quotes) GetAllQuotesByCategory(category string) ([]model.Quote, error) {
	quote.Lock()
	defer  quote.Unlock()
	quotes := []model.Quote{}

	for _, value := range quote.Quotes {
		if value.Category == category {
			quotes = append(quotes, value)

		}
	}
	if quotes == nil {
		return nil, constants.ErrNotFound
	}

	return quotes, nil
}