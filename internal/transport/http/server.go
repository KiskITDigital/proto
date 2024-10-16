package http

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/config"
)

type Server struct {
	srv    *http.Server
	logger *slog.Logger
}

func NewServer(logger *slog.Logger, cfg config.HTTP, router *Router) (*Server, error) {
	apiServer, err := api.NewServer(
		router,
		api.WithErrorHandler(router.HandleError),
	)
	if err != nil {
		return nil, fmt.Errorf("init http server: %w", err)
	}

	mux := chi.NewMux()
	mux.Mount("/", apiServer)

	registerSwaggerUIHandlers(logger, mux, cfg.SwaggerUIPath)

	server := &Server{
		srv: &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.Port),
			Handler: mux,
		},
		logger: logger,
	}

	return server, nil
}

func (s *Server) Start() error {
	s.logger.Info("Starting http server", "addr", s.srv.Addr)
	err := s.srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) Stop() error {
	s.logger.Info("Shutting down http server")
	return s.srv.Shutdown(context.Background())
}
