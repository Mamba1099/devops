package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsResponse struct {
	Success bool `json:"success"`
	Metrics struct {
		Users   int     `json:"users"`
		Uptime  float64 `json:"uptime"`
		Memory  uint64  `json:"memory_mb"`
		CPU     int     `json:"cpu_cores"`
		GoRoutines int `json:"goroutines"`
	} `json:"metrics"`
}

func MetricsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users int
		db.QueryRow("SELECT COUNT(*) FROM users").Scan(&users)

		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		response := MetricsResponse{
			Success: true,
		}
		response.Metrics.Users = users
		response.Metrics.Uptime = time.Since(startTime).Seconds()
		response.Metrics.Memory = m.Alloc / 1024 / 1024
		response.Metrics.CPU = runtime.NumCPU()
		response.Metrics.GoRoutines = runtime.NumGoroutine()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func PrometheusHandler() http.Handler {
	return promhttp.Handler()
}