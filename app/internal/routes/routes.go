package routes

import (
	"database/sql"
	"net/http"

	"devops-template/internal/handlers"
	"devops-template/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	userHandler := handlers.NewUserHandler(db)
	router.HandleFunc("/health", handlers.HealthHandler).Methods("GET")
	router.HandleFunc("/ready", handlers.ReadyHandler(func() bool {
		return db.Ping() == nil
	})).Methods("GET")

	router.HandleFunc("/metrics", middleware.MetricsMiddleware(handlers.MetricsHandler(db))).Methods("GET")
	router.Handle("/metrics/prometheus", handlers.PrometheusHandler()).Methods("GET")
	router.HandleFunc("/api/users", middleware.MetricsMiddleware(userHandler.GetAll)).Methods("GET")
	router.HandleFunc("/api/users", middleware.MetricsMiddleware(userHandler.Create)).Methods("POST")
	router.HandleFunc("/api/users/{id}", middleware.MetricsMiddleware(userHandler.GetByID)).Methods("GET")
	router.HandleFunc("/api/users/{id}", middleware.MetricsMiddleware(userHandler.Update)).Methods("PUT")
	router.HandleFunc("/api/users/{id}", middleware.MetricsMiddleware(userHandler.Delete)).Methods("DELETE")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"message": "DevOps Template API",
			"version": "1.0.0",
			"endpoints": {
				"health": "/health",
				"ready": "/ready",
				"users": "/api/users",
				"metrics": "/metrics",
				"prometheus": "/metrics/prometheus"
			}
		}`))
	}).Methods("GET")

	return router
}