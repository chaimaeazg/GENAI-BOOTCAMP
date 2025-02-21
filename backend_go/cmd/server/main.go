package main

import (
	"backend_go/internal/handlers"
	"backend_go/internal/models"
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"backend_go/internal/config"
	"backend_go/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg := config.Load()

	// Initialize DB with config
	db, err := models.InitDB(cfg.DBPath)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	// Setup logging
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatal("Failed to create logs directory:", err)
	}
	logFile, err := os.Create("logs/server.log")
	if err != nil {
		log.Fatal("Could not create log file:", err)
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	// Create a new Gin router
	router := gin.Default()

	// Add middleware
	router.Use(middleware.RequestLogger())

	// Add database to context
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Define a simple health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	// Placeholder for API endpoints
	router.GET("/api/dashboard/last_study_session", getLastStudySession)
	router.GET("/api/dashboard/study_progress", handlers.GetStudyProgress)

	// Define additional API endpoints
	router.GET("/api/study_activities/:id", handlers.GetStudyActivity)
	router.GET("/api/study_activities/:id/success_rate", handlers.GetActivitySuccessRate)
	router.POST("/api/study_activities/:id/launch", handlers.LaunchStudyActivity)

	// Groups endpoints
	router.GET("/api/groups", handlers.GetAllGroups)
	router.GET("/api/groups/:id", handlers.GetGroupByID)
	router.GET("/api/groups/:id/words", handlers.GetGroupWords)
	router.GET("/api/groups/:id/study_sessions", handlers.GetGroupStudySessions)

	// Study session endpoints
	router.POST("/api/study_sessions", handlers.CreateStudySession)
	router.POST("/api/study_sessions/:id/words/:word_id/review", handlers.RecordReview)

	// System endpoints
	router.POST("/api/reset_history", handlers.ResetHistory)
	router.POST("/api/full_reset", handlers.FullReset)
	router.POST("/api/import-initial-data", handlers.ImportInitialData)

	// Study Sessions
	router.GET("/api/study_sessions", handlers.GetAllStudySessions)
	router.GET("/api/study_sessions/:id", handlers.GetStudySessionDetails)
	router.GET("/api/study_sessions/:id/words", handlers.GetSessionWords)

	// Study Activities
	router.GET("/api/study_activities", handlers.GetAllStudyActivities)
	router.POST("/api/study_activities", handlers.CreateStudyActivity)
	router.GET("/api/study_activities/:id/study_sessions", handlers.GetActivityStudySessions)

	// Words
	router.GET("/api/words", handlers.GetWords)
	router.GET("/api/words/:id", handlers.GetWordDetails)
	router.GET("/api/words/:id/study_sessions", handlers.GetWordStudySessions)

	// Dashboard
	router.GET("/api/dashboard/quick-stats", handlers.GetQuickStats)

	// Start the server on port 8080
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()

	// Wait for interrupt
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Forced shutdown:", err)
	}
	log.Println("Server exited")
}

// Define your handler functions here
// Example:
func getLastStudySession(c *gin.Context) {
	// Logic to retrieve the last study session
	c.JSON(http.StatusOK, gin.H{
		"id":                456,
		"group_id":          789,
		"created_at":        "2024-03-15T09:45:12-05:00",
		"study_activity_id": 234,
		"group_name":        "Common Verbs",
	})
}
