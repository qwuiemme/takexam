package server

import (
	"net/http"
	"time"
)

type Server struct {
	engine *http.Server
}

func (s *Server) Run(address string) error {
	s.engine = &http.Server{
		Addr:        address,
		Handler:     DetermineRoutes(),
		ReadTimeout: 10 * time.Second,
	}

	return s.engine.ListenAndServe()
}
