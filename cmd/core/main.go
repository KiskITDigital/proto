package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"gitlab.ubrato.ru/ubrato/core/internal/config"
	dadataGateway "gitlab.ubrato.ru/ubrato/core/internal/gateway/dadata"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/token"
	authService "gitlab.ubrato.ru/ubrato/core/internal/service/auth"
	catalogService "gitlab.ubrato.ru/ubrato/core/internal/service/catalog"
	tenderService "gitlab.ubrato.ru/ubrato/core/internal/service/tender"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"gitlab.ubrato.ru/ubrato/core/internal/store/postgres"
	"gitlab.ubrato.ru/ubrato/core/internal/transport/http"
	authHandler "gitlab.ubrato.ru/ubrato/core/internal/transport/http/handlers/auth"
	catalogHandler "gitlab.ubrato.ru/ubrato/core/internal/transport/http/handlers/catalog"
	errorHandler "gitlab.ubrato.ru/ubrato/core/internal/transport/http/handlers/error"
	tendersHandler "gitlab.ubrato.ru/ubrato/core/internal/transport/http/handlers/tenders"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg, err := config.ReadConfig()
	if err != nil {
		logger.Error("Error parsing default config from env", "error", err)
		os.Exit(1)
	}

	if cfg.Debug {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))

		logger.Warn("Debug enabled")
	}

	if err := run(cfg, logger); err != nil {
		logger.Error("Error initializing service", "error", err)
		os.Exit(1)
	}
}

func run(cfg config.Default, logger *slog.Logger) error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	psqlDB, err := postgres.New(cfg.Store.Postgres)
	if err != nil {
		return fmt.Errorf("create postgres: %w", err)
	}
	psql := store.New(psqlDB)

	userStore := postgres.NewUserStore()
	organizationStore := postgres.NewOrganizationStore()
	sessionStore := postgres.NewSessionStore()
	tenderStore := postgres.NewTenderStore()
	catalogStore := postgres.NewCatalogStore()

	dadataGateway := dadataGateway.NewClient(cfg.Gateway.Dadata.APIKey)

	tokenAuthorizer, err := token.NewTokenAuthorizer(cfg.Auth.JWT)
	if err != nil {
		return fmt.Errorf("init token authorizer: %w", err)
	}

	authService := authService.New(
		psql,
		userStore,
		organizationStore,
		sessionStore,
		dadataGateway,
		tokenAuthorizer,
	)

	tenderService := tenderService.New(
		psql,
		tenderStore,
	)

	catalogService := catalogService.New(
		psql,
		catalogStore,
	)

	router := http.NewRouter(http.RouterParams{
		Error:   errorHandler.New(logger),
		Auth:    authHandler.New(logger, authService),
		Tenders: tendersHandler.New(logger, tenderService),
		Catalog: catalogHandler.New(logger, catalogService),
	})

	server, err := http.NewServer(logger, cfg.Transport.HTTP, router)
	if err != nil {
		return fmt.Errorf("create http server: %w", err)
	}

	go func() {
		<-sig
		logger.Info("Received termination signal, cleaning up")
		err := server.Stop()
		if err != nil {
			logger.Error("Stop http server", "error", err)
		}
	}()

	err = server.Start()
	if err != nil {
		return fmt.Errorf("serve http: %w", err)
	}

	return nil
}
