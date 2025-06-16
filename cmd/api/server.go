package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	router *chi.Mux
}

func NewServer() Server {

	server := Server{}

	server.mountRoutes()
	return server
}

func (s *Server) mountRoutes() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/health", checkHealthHandler)

	s.router = r
}

func (s *Server) Start(address string) error {
	return http.ListenAndServe(address, s.router)
}
