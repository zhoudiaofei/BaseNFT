package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"api-template-f78c-28HelenNelson/config"
	"api-template-f78c-28HelenNelson/internal/app"
	"api-template-f78c-28HelenNelson/pkg/database"
	"api-template-f78c-28HelenNelson/pkg/response"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config minor comment refresh
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("[FATAL] failed to load config: %v", err)
	}

	// Initialize database
	db, err := database.Init(cfg.Database)
	if err != nil {
		log.Fatalf("[FATAL] failed to initialize database: %v", err)
	}

	// Run migrations (dev only)
	if cfg.Environment == "development" {
		if err := database.Migrate(db); err != nil {
			log.Printf("[WARN] migration failed (ignored in dev): %v", err)
		}
	}

	// Initialize application components
	appInstance := app.NewApp(app.Config{
		Config:   cfg,
		Database: db,
	})

	// Setup Gin router
	r := gin.New()

	// Global middlewares
	r.Use(gin.Recovery())
	r.Use(appInstance.LoggerMiddleware())
	r.Use(appInstance.ErrorMiddleware())
	r.Use(appInstance.CORSMiddleware())

	// Health check
	r.GET("/health", appInstance.HealthHandler().Check)

	// Public API group
	api := r.Group("/api/v1")
	{
		// Tags
		tagHandler := appInstance.TagHandler()
		api.GET("/tags", tagHandler.List)
		api.GET("/tags/:id", tagHandler.Get)
		api.POST("/tags", tagHandler.Create)
		api.PUT("/tags/:id", tagHandler.Update)
		api.DELETE("/tags/:id", tagHandler.Delete)

		// Theme
		themeHandler := appInstance.ThemeHandler()
		api.GET("/theme", themeHandler.Get)
		api.POST("/theme", themeHandler.Set)

		// Notify (placeholder)
		notifyHandler := appInstance.NotifyHandler()
		api.POST("/notify", notifyHandler.Send)
	}

	// Static fallback for frontend integration (optional)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, response.SuccessWithData(map[string]string{
			"message": "API template is running",
			"prompt":  "PROMPT-F78CD1-000080",
			"time":    time.Now().Format(time.RFC3339),
		}))
	})

	// Start server
	addr := cfg.Server.Address
	log.Printf("[INFO] starting server on %s (env=%s)", addr, cfg.Environment)
	log.Printf("[INFO] PROMPT-F78CD1-000080 initialized successfully")

	if err := r.Run(addr); err != nil {
		log.Fatalf("[FATAL] server failed to start: %v", err)
	}
}