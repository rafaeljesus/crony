package handlers

import (
	"net/http"

	"github.com/EmpregoLigado/cron-srv/pkg/checker"
	"github.com/EmpregoLigado/cron-srv/pkg/response"
)

type HealthzHandler struct {
	checkers map[string]checker.Checker
}

func NewHealthzHandler(checkers map[string]checker.Checker) *HealthzHandler {
	return &HealthzHandler{checkers}
}

func (h *HealthzHandler) HealthzIndex(w http.ResponseWriter, r *http.Request) {
	payload := make(map[string]bool)

	for k, v := range h.checkers {
		payload[k] = v.IsAlive()
	}

	response.JSON(w, http.StatusOK, payload)
}
