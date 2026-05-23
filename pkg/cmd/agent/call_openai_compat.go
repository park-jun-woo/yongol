//ff:func feature=agent type=adapter control=sequence
//ff:what callOpenAICompat — OpenAI 호환 endpoint (xai) chat completion 호출

package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// callOpenAICompat calls an OpenAI-compatible chat completions endpoint (xai).
func callOpenAICompat(endpoint, backend, model, system, user string) (string, error) {
	apiKey, err := loadAPIKey(backend)
	if err != nil {
		return "", err
	}
	body := map[string]any{
		"model": model,
		"messages": []map[string]string{
			{"role": "system", "content": system},
			{"role": "user", "content": user},
		},
		"max_tokens":  2048,
		"temperature": 0,
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("marshal %s request: %w", backend, err)
	}
	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("create %s request: %w", backend, err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("%s request: %w", backend, err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read %s response: %w", backend, err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%s %d: %s", backend, resp.StatusCode, string(data))
	}
	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return "", fmt.Errorf("parse %s response: %w", backend, err)
	}
	if len(result.Choices) == 0 {
		return "", fmt.Errorf("%s: empty choices", backend)
	}
	return result.Choices[0].Message.Content, nil
}
