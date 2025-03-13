package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"nihal-innsof/portfolio/internal/views"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	logger *log.Logger
	port   int
	server *http.Server
}

func NewServer(logger *log.Logger, port int) (*server, error) {
	if logger == nil {
		return nil, fmt.Errorf("logger is required")
	}

	return &server{
		logger: logger,
		port:   port,
	}, nil
}

func (s *server) Start() error {
	s.logger.Printf("Starting server on port %d", s.port)
	var stopChan chan os.Signal

	router := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	router.HandleFunc("GET /", s.homeHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: router,
	}

	stopChan = make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error when starting server: %v", err)
		}
	}()

	<-stopChan

	if err := s.server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Error when shutting down server: %v", err)
		return err
	}

	return nil
}

func (s *server) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	homeTemplate := views.HomeView()

	homeTemplate.Render(r.Context(), w)
}
