package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"apps-gin/interceptors"
	logger "apps-gin/internal"
	"apps-gin/internal/telemetry"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
	baseCtx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No .env file found")
	}

	logger.Init()
	defer logger.Sync()

	shutdown := telemetry.Init("apps-gin")
	defer shutdown(baseCtx)

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	// open telemetry middleware
	router.Use(otelgin.Middleware("apps-gin"))
	// catch panics
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	protected := router.Group("/api")
	protected.Use(interceptors.JWTAuthMiddleware())
	{
		protected.GET("/profiles", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"profiles": []interface{}{}})
		})
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Println("Server started on :8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutdown signal received")

	// context timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown:", err)
	}

	log.Println("Server exiting!")
}
