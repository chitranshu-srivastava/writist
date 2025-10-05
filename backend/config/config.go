package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	LLMProvider    string // "claude", "openai", or "gemini"
	ClaudeAPIKey   string
	OpenAIAPIKey   string
	GeminiAPIKey   string
	ModelName      string // Model to use for selected provider
	ClaudeModel    string
	Port           string
	Environment    string
	AllowedOrigins []string
	RateLimit      int
	MaxTextLength  int
}

func Load() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}
	// parse allowed origins
	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")

	for i, origin := range allowedOrigins {
		allowedOrigins[i] = strings.TrimSpace(origin)
	}

	// Prase rate limit
	rateLimit, err := strconv.Atoi(getEnv("RATE_LIMIT", "30"))
	if err != nil {
		rateLimit = 30
	}

	maxTextLength, err := strconv.Atoi(getEnv("MAX_TEXT_LENGTH", "5000"))
	if err != nil {
		maxTextLength = 5000
	}

	config := &Config{
		LLMProvider:    getEnv("LLM_PROVIDER", "claude"),
		ClaudeAPIKey:   os.Getenv("CLAUDE_API_KEY"),
		OpenAIAPIKey:   os.Getenv("OPENAI_API_KEY"),
		GeminiAPIKey:   os.Getenv("GEMINI_API_KEY"),
		ModelName:      getEnv("MODEL_NAME", ""),
		ClaudeModel:    getEnv("CLAUDE_MODEL", "claude-sonnet-4-20250514"),
		Port:           getEnv("PORT", "8080"),
		Environment:    getEnv("ENVIRONMENT", "development"),
		AllowedOrigins: allowedOrigins,
		RateLimit:      rateLimit,
		MaxTextLength:  maxTextLength,
	}
	// Validate required fields
	if config.ClaudeAPIKey == "" {
		log.Fatal("CLAUDE_API_KEY environment variable is required")
	}

	return config
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
