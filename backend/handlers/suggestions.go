package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/chitranshu-srivastava/writist/backend/services"
)

type SuggestionsHandler struct {
	llmService    services.LLMService
	maxTextLength int
}

func NewSuggestionsHandler(llmService services.LLMService, maxTextLength int) *SuggestionsHandler {
	return &SuggestionsHandler{
		llmService:    llmService,
		maxTextLength: maxTextLength,
	}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (h *SuggestionsHandler) HandleSuggestions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req services.SuggestionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate text
	req.Text = strings.TrimSpace(req.Text)
	if req.Text == "" {
		h.sendError(w, "Text is required", http.StatusBadRequest)
		return
	}

	if len(req.Text) > h.maxTextLength {
		h.sendError(w, "Text is too long", http.StatusBadRequest)
		return
	}

	if req.Type == "" {
		req.Type = "general"
	}

	// Get suggestions
	suggestions, err := h.llmService.GetSuggestions(req)
	if err != nil {
		h.sendError(w, "Failed to get suggestions: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestions)
}

func (h *SuggestionsHandler) HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"service": "writist",
	})
}

func (h *SuggestionsHandler) sendError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
