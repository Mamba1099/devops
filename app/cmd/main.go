package main

import (
	"log"
	"net/http"
	"os"

	"devops-template/internal/config"
	"devops-template/internal/database"
	"devops-template/internal/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	cfg := config.Load()

	db, err := database.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := database.Migrate(db); err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}

	if err := database.Seed(db); err != nil {
		log.Printf("⚠️ Seed warning: %v", err)
	}

	router := routes.SetupRoutes(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Server running on port %s", port)
	log.Printf("📊 Health: http://localhost:%s/health", port)
	log.Printf("📊 Metrics: http://localhost:%s/metrics", port)
	log.Printf("📊 API: http://localhost:%s/api/users", port)

	log.Printf("✅ Server started successfully")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}