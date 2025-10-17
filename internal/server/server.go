package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	logger *log.Logger
	server *http.Server
}

func MyServer(logger *log.Logger) *Server {
	// создаем http-роутер
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.OutputHTMLHandler)
	router.HandleFunc("/upload", handlers.UpLoadHandler)

	httpMyServ := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &Server{
		server: httpMyServ,
		logger: logger,
	}
}

func (s *Server) ListenAndServe() error {
	return s.server.ListenAndServe()
}
