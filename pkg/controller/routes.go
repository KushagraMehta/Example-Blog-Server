package controller

import middlewares "github.com/KushagraMehta/Example-Blog-Server/pkg/middleware"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home))
}
