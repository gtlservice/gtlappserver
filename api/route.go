package api

func (s *ApiServer) InitRouter() {

	s.echo.Get("/service1/v1/catalog", s.getCatalogHandleFunc)
	s.echo.Post("/service1/v1/user", s.postUserHandleFunc)
}
