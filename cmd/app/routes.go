package app


func (server *server) Init(){
	server.router.POST("/api/quote/create", server.handleCreateQuote)
	server.router.POST("/api/quote/editquote", server.handlerEditQuote)
	server.router.DELETE("/api/quote/delete/:id", server.handleDeleteQuote)
	server.router.GET("/api/quote/quotes", server.handlerGetAllQuotes)
	server.router.GET("/api/quote/quotes/:categorytype", server.handleGetAllQuotesByCategory)
	server.router.GET("/api/quote/quotes/random", server.handleGetRandomQuote)
}