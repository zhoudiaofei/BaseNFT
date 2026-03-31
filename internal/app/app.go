package app

import (
	"api-template-f78c-28HelenNelson/config"
	"api-template-f78c-28HelenNelson/internal/handler"
	"api-template-f78c-28HelenNelson/internal/middleware"
	"api-template-f78c-28HelenNelson/pkg/database"
	"api-template-f78c-28HelenNelson/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// App holds the core application components minor comment refresh
type App struct {
	Router *gin.Engine
	DB     *gorm.DB
	Config *config.Config
}

// NewApp initializes and returns a new App instance
func NewApp(cfg *config.Config) (*App, error) {
	// Initialize Gin
	gin.SetMode(gin.ReleaseMode)
	if cfg.App.Debug {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	// Load middleware
	r.Use(middleware.Logger())
	r.Use(middleware.ErrorHandler())
	r.Use(middleware.CORS())

	// Initialize database
	db, err := database.Init(cfg.Database)
	if err != nil {
		return nil, err
	}

	// Run migrations in dev mode only
	if cfg.App.Debug {
		if err := database.Migrate(db); err != nil {
			return nil, err
		}
	}

	// Initialize handlers with dependencies
	handlers := handler.NewHandler(db, cfg)

	// Register routes
	handlers.RegisterRoutes(r)

	return &App{
		Router: r,
		DB:     db,
		Config: cfg,
	}, nil
}

// Run starts the HTTP server
func (a *App) Run() error {
	addr := a.Config.Server.Addr
	return a.Router.Run(addr)
}

// Shutdown gracefully closes resources (placeholder for future extension)
func (a *App) Shutdown() error {
	// e.g., close DB connection, stop background jobs
	return nil
}