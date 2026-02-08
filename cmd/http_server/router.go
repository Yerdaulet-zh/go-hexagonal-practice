package httpserver

import (
	"net/http"
	"time"

	"github.com/go-hexagonal-practice/internal/adapters/handlers/http/middleware"
	"github.com/go-hexagonal-practice/internal/adapters/repository/postgre"
	"github.com/go-hexagonal-practice/internal/core/ports"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"
)

type Middleware func(http.Handler) http.Handler

// Create a helper to chain them
func ApplyMiddleware(h http.Handler, mws ...Middleware) http.Handler {
	for _, mw := range mws {
		h = mw(h)
	}
	return h
}

func MapBusinessRoutes(logger ports.Logger, rdb *redis.Client) http.Handler {
	mux := http.NewServeMux()

	middlewares := []Middleware{
		middleware.LoggingMiddleware(logger),                   // 3. Log everything (including blocks)
		middleware.IPRateLimiter(logger, rdb, 10, time.Minute), // 2. Then check limit
		// middleware.RecoveryMiddleware(logger),               // 1. Catch panics first
	}
	return ApplyMiddleware(mux, middlewares...)
}

func MapManagementRoutes(logger ports.Logger, db *postgre.Client, reg *prometheus.Registry) http.Handler {
	mux := http.NewServeMux()

	healthHdl := NewHealthHandler(db)
	mux.HandleFunc("GET /healthz", healthHdl.Healthz)
	mux.HandleFunc("GET /ready", healthHdl.Ready)

	mux.Handle("GET /metrics", promhttp.Handler())
	return mux
}
