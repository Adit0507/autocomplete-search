package api

import (
	"encoding/json"
	"net/http"

	"github.com/Adit0507/autocomplete-search/config"
	"github.com/Adit0507/autocomplete-search/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.AutocompleteService
}

func NewRouter() *mux.Router {
	cfg := config.NewConfig()
	handler := &Handler{
		service: service.NewAutoCompleteService(cfg.CacheSize),
	}

	router := mux.NewRouter()
	router.HandleFunc("/autocomplete", handler.AutoCompleteHandler)

	return router
}

func (h *Handler) AutoCompleteHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	suggestions := h.service.GetSuggestions(query, 5)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(suggestions); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
