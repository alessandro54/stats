package common

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
)

type ServerConfig struct {
	Port         int
	App          *fiber.App
	GracefulWait time.Duration
}

func StartFiberServer(cfg ServerConfig) error {
	addr := fmt.Sprintf(":%d", cfg.Port)

	// Start server in a goroutine
	go func() {
		log.Printf("🚀 Fiber server running on %s", addr)
		if err := cfg.App.Listen(addr); err != nil {
			log.Fatalf("❌ Failed to start Fiber server: %v", err)
		}
	}()

	// Listen for OS signals for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down Fiber server...")

	// Graceful shutdown
	time.Sleep(cfg.GracefulWait) // optional: give requests time to finish

	if err := cfg.App.Shutdown(); err != nil {
		return fmt.Errorf("failed to shutdown Fiber: %w", err)
	}

	if err := cfg.App.Shutdown(); err != nil {
		return fmt.Errorf("failed to shutdown Fiber: %w", err)
	}

	log.Println("✅ Fiber server exited cleanly")
	return nil
}
