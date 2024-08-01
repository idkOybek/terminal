package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/idkOybek/newNewTerminal/docs"
	"github.com/idkOybek/newNewTerminal/internal/config"
	"github.com/idkOybek/newNewTerminal/internal/handler"
	customMiddleware "github.com/idkOybek/newNewTerminal/internal/middleware"
	"github.com/idkOybek/newNewTerminal/internal/repository"
	"github.com/idkOybek/newNewTerminal/internal/service"
	"github.com/idkOybek/newNewTerminal/pkg/database"
	"github.com/idkOybek/newNewTerminal/pkg/logger"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/idkOybek/newNewTerminal/internal/repository/postgres"

	"go.uber.org/zap"
)

// @title Your API Title
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host newnewterminal.onrender.com
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// Initialize logger
	logger := logger.NewLogger()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load config", zap.Error(err))
	}

	// Connect to database
	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	defer db.Close()

	// Initialize repositories
	repos := repository.NewRepositories(db)

	// Initialize services
	services := service.NewServices(service.Deps{
		Repos: repos,
	})

	// Initialize handlers
	authHandler := handler.NewAuthHandler(services.Auth, logger)
	userHandler := handler.NewUserHandler(services.User, logger)
	fiscalModuleHandler := handler.NewFiscalModuleHandler(services.FiscalModule, logger)
	terminalHandler := handler.NewTerminalHandler(services.Terminal, logger)

	// Set up router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(customMiddleware.LoggerMiddleware(logger))

	// Swagger
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	// Routes
	r.Route("/api", func(r chi.Router) {
		r.Mount("/auth", authHandler.Routes())
		r.Mount("/users", userHandler.Routes())
		r.Group(func(r chi.Router) {
			r.Use(customMiddleware.AuthMiddleware(logger))
			r.Mount("/fiscal-modules", fiscalModuleHandler.Routes())
			r.Mount("/terminals", terminalHandler.Routes())
		})
	})

	// Set up server
	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: r,
	}

	// Start server
	go func() {
		logger.Info("Starting server", zap.String("port", cfg.ServerPort))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	logger.Info("Server exiting")
}
