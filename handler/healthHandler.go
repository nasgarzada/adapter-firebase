package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
)

// HealthHandler is for handler usage in main.go
type HealthHandler struct{}

// NewHealthHandler for kubernetes health and readiness check
func NewHealthHandler(router *chi.Mux) *HealthHandler {
	h := &HealthHandler{}

	router.Get("/readiness", h.Health)
	router.Get("/health", h.Health)
	router.Handle("/metrics", promhttp.Handler())
	router.Get("/swagger/*", httpSwagger.Handler())
	return h
}

// @Summary Health endpoint for kubernetes health and readiness check
// @Tags health-handler
// @Success 200 {} http.Response
// @Router /health [get]
func (*HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
