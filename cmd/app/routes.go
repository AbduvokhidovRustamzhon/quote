package app


func (s *server) Init(){
	s.router.POST("/api/quote/create", s.handleCreateQuote)
	s.router.POST("/api/quote/editquote", s.handlerEditQuote)
}