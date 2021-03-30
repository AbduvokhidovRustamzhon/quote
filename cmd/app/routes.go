package app

func (s *server) Init() {
	s.router.POST("/api/quotes", s.createQuote)
	s.router.PUT("/api/quotes/:id", s.editQuote)
	s.router.DELETE("/api/quotes/:id", s.deleteQuote)
	s.router.GET("/api/quotes", s.getAllQuotes)
	s.router.GET("/api/quotes/:categorytype", s.getAllQuotesByCategory)
	s.router.GET("/api/random/quotes/", s.handleGetRandomQuote)
}
