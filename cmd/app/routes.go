package app


func (s *server) Init(){
	s.router.POST("/api/quote/create", s.handleCreateQuote)
	s.router.POST("/api/quote/editquote", s.handlerEditQuote)
	s.router.DELETE("/api/quote/delete/:id", s.handleDeleteQuote)
	s.router.GET("/api/quote/quotes", s.handlerGetAllQuotes)
}