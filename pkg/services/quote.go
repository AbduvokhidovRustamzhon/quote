package services

import (
	"sync"

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
func (q *Quotes) CreateQuote(quote *model.Quote) (err error) {
	q.Lock()
	defer  q.Unlock()
	quote.ID = uuid.New().String()
	q.Quotes[quote.ID] = *quote

	if q.Quotes == nil {
		return err
	}
	return nil
}