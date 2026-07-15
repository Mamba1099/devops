package handlers

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

type HealthResponse struct {
	Status    string            `json:"status"`
	Timestamp string            `json:"timestamp"`
	Uptime    float64           `json:"uptime"`
	Memory    map[string]uint64 `json:"memory"`
	CPU       int               `json:"cpu"`
	Version   string            `json:"version"`
}

type ReadyResponse struct {
	Status    string `json:"status"`
	Database  string `json:"database"`
	Timestamp string `json:"timestamp"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Uptime:    time.Since(startTime).Seconds(),
		Memory: map[string]uint64{
			"alloc":       m.Alloc / 1024 / 1024,
			"total_alloc": m.TotalAlloc / 1024 / 1024,
			"sys":         m.Sys / 1024 / 1024,
			"num_gc":      uint64(m.NumGC),
		},
		CPU:     runtime.NumCPU(),
		Version: runtime.Version(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

var startTime = time.Now()

func ReadyHandler(dbReady func() bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := "ready"
		dbStatus := "connected"

		if !dbReady() {
			status = "not ready"
			dbStatus = "disconnected"
		}

		response := ReadyResponse{
			Status:    status,
			Database:  dbStatus,
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		}

		if status == "ready" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		json.NewEncoder(w).Encode(response)
	}
}