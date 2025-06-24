package api_services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"lang-learner-wails/constants/url_constants"
	"net/http"
	"os"
	"time"
)

func ElevenLabsTextToSpeechAudioBytes(text string) ([]byte, error) {
	apiKey := os.Getenv("ELEVEN_LABS_API_KEY")
	println("API Key: " + apiKey)
	// Prepare request body
	bodyData := map[string]string{
		"text":     text,
		"model_id": "eleven_multilingual_v2",
	}
	jsonBody, err := json.Marshal(bodyData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", url_constants.ElevenLabsURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add headers
	req.Header.Set("xi-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	fmt.Println("Making API call.")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	// Check status
	if resp.StatusCode == http.StatusOK {
		audioData, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read audio response: %w", err)
		}
		return audioData, nil
	}

	// Not 200 OK
	errMsg, _ := io.ReadAll(resp.Body)
	return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, errMsg)
}
