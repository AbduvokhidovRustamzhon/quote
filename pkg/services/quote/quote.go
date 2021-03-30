package quote

import (
	"github.com/AbduvokhidovRustamzhon/quote/pkg/services/quote/models"
	"github.com/AbduvokhidovRustamzhon/quote/types"
	"math/rand"
	"sync"
	"time"
)

type Quotes struct {
	sync.Mutex
	Quotes map[string]*models.Quote `json:"quotes"`
	rand   *rand.Rand
}

func NewQuotes() *Quotes {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return &Quotes{
		Quotes: make(map[string]*models.Quote),
		rand:   r,
	}
}

// Create Quotes
func (q *Quotes) CreateQuote(quote *models.Quote) error {
	q.Lock()
	defer q.Unlock()

	q.Quotes[quote.ID] = quote
	return nil
}

// Edit Quote 
func (q *Quotes) EditQuote(id string, quoteRequest *models.Request) error {
	q.Lock()
	defer q.Unlock()

	quote, ok := q.Quotes[id]
	if !ok {
		return types.ErrNotFound
	}
	quote.Author = quoteRequest.Author
	quote.Quote = quoteRequest.Quote
	quote.Category = quoteRequest.Category
	return nil
}

// Delete Quote
func (q *Quotes) Delete(id string) error {
	_, exists := q.Quotes[id]
	if !exists {
		return types.ErrNotFound
	}
	delete(q.Quotes, id)
	return nil
}

// Get Quotes
func (q *Quotes) GetAllQuotes() ([]models.Quote, error) {
	q.Lock()
	defer q.Unlock()

	quotes := make([]models.Quote, len(q.Quotes))
	for _, quote := range q.Quotes {
		quotes = append(quotes, *quote)
	}
	return quotes, nil
}

// Get All Quotes by Category
func (q *Quotes) GetAllQuotesByCategory(category string) ([]models.Quote, error) {
	q.Lock()
	defer q.Unlock()

	quotes := make([]models.Quote, len(q.Quotes))
	for _, quote := range q.Quotes {
		if quote.Category == category {
			quotes = append(quotes, *quote)
		}
	}
	return quotes, nil
}

// Get Quote from randomizer
func (q *Quotes) GetRandomQuote() (*models.Quote, error) {
	q.Lock()
	defer q.Unlock()

	length := len(q.Quotes)
	if length == 0 {
		return nil, types.ErrNotFound
	}

	randomNumber := q.rand.Intn(length)
	count := 0
	var randomQuote models.Quote
	for _, quote := range q.Quotes {
		if count == randomNumber {
			randomQuote = *quote
			break
		}
		count++
	}
	return &randomQuote, nil
}

// DeleteOldQuotes deletes old quotes that were created 'period' ago
func (q *Quotes) DeleteOldQuotes(period time.Duration) func() {
	return func() {
		now := time.Now().Add(-period)
		for _, quote := range q.Quotes {
			if now.After(quote.CreatedAt) {
				_ = q.Delete(quote.ID)
			}
		}
	}
}

// Check validity for time passed
func timePassed(check, date time.Time) bool {
	return check.After(date)
}

// worker that will check and delete old quotes
func Worker(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}
