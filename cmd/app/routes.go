package app


func (s *server) Init(){
	s.router.POST("/api/quote/create", s.handleCreateQuote)
}