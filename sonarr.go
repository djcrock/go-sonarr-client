package sonarr

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// Sonarr contains fields needed to make API calls to a Sonarr server
type Sonarr struct {
	baseURL    url.URL
	apiKey     string
	HTTPClient http.Client
}

// New creates a new Sonarr client instance
func New(apiURL, apiKey string) (*Sonarr, error) {
	if apiURL == "" {
		return &Sonarr{}, errors.New("apiURL is required")
	}

	if apiKey == "" {
		return &Sonarr{}, errors.New("apiKey is required")
	}

	baseURL, err := url.Parse(apiURL)
	if err != nil {
		return &Sonarr{}, fmt.Errorf("Failed to parse baseURL: %v", err)
	}

	return &Sonarr{
		baseURL:    baseURL,
		apiKey:     apiKey,
		HTTPClient: http.Client{},
	}, nil
}
