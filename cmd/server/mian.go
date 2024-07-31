// cmd/server/main.go

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/idkOybek/newNewTerminal/internal/config"
	"github.com/idkOybek/newNewTerminal/internal/handler"
	"github.com/idkOybek/newNewTerminal/internal/repository/postgres"
	"github.com/idkOybek/newNewTerminal/internal/service"
	"github.com/idkOybek/newNewTerminal/pkg/database"
	"github.com/idkOybek/newNewTerminal/pkg/logger"

	ware "github.com/idkOybek/newNewTerminal/internal/middleware/logger_middleware.go"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Initialize configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	logger.Init(cfg.LogLevel)

	// Connect to database
	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	userRepo := postgres.NewUserRepository(db)
	fiscalModuleRepo := postgres.NewFiscalModuleRepository(db)
	terminalRepo := postgres.NewTerminalRepository(db)
	linkRepo := postgres.NewLinkRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)
	fiscalModuleService := service.NewFiscalModuleService(fiscalModuleRepo)
	terminalService := service.NewTerminalService(terminalRepo, linkRepo)
	linkService := service.NewLinkService(linkRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	fiscalModuleHandler := handler.NewFiscalModuleHandler(fiscalModuleService)
	terminalHandler := handler.NewTerminalHandler(terminalService)
	linkHandler := handler.NewLinkHandler(linkService)

	// Initialize router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(ware.LoggerMiddleware)

	// Routes
	r.Route("/api", func(r chi.Router) {
		r.Mount("/users", userHandler.Routes())
		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)
			r.Mount("/fiscal-modules", fiscalModuleHandler.Routes())
			r.Mount("/terminals", terminalHandler.Routes())
			r.Mount("/links", linkHandler.Routes())
		})
	})

	// Start server
	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: r,
	}

	// Graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Failed to start server: %v", err)
		}
	}()

	logger.Infof("Server started on port %s", cfg.ServerPort)

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.Info("Server exited properly")
}
