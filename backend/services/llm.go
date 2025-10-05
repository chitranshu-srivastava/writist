package services

type LLMService interface {
	GetSuggestions(req SuggestionRequest) (*SuggestionRequest, error)
}

type SuggestionRequest struct {
	Text string `json:"text"`
	Type string `json:"type"` // "grammar", "style", "tone", "general"
}

type SuggestionResponse struct {
	Suggestions []Suggestion `json:"suggestions"`
	Original    string       `json:"original"`
}

type Suggestion struct {
	Type        string    `json:"type"`
	Issue       string    `json:"issue"`
	Suggestion  string    `json:"suggestion"`
	Replacement string    `json:"replacement,omitempty"`
	Position    *Position `json:"position,omitempty"`
}

type Position struct {
	Start int `json:"start"`
	End   int `json:"end"`
}
