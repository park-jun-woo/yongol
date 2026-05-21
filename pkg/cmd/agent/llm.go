//ff:func feature=agent type=adapter control=sequence
//ff:what llmCall — LLM 호출 추상화 (ollama/xai/gemini)

package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// llmCall sends a chat completion request to the configured backend and returns
// the assistant reply content.
func llmCall(backend, model, systemPrompt, userPrompt string) (string, error) {
	switch backend {
	case "ollama":
		return callOllama(model, systemPrompt, userPrompt)
	case "xai":
		return callOpenAICompat("https://api.x.ai/v1/chat/completions", backend, model, systemPrompt, userPrompt)
	case "gemini":
		return callGemini(model, systemPrompt, userPrompt)
	default:
		return "", fmt.Errorf("unsupported backend: %s", backend)
	}
}

// LLMCallFunc is the signature for LLM call functions used by scaffold.
type LLMCallFunc func(backend, model, system, user string) (string, error)

// llmCallWithNumCtx calls the LLM with a custom num_ctx for ollama.
// For non-ollama backends, numCtx is ignored and the standard call is used.
func llmCallWithNumCtx(backend, model, system, user string, numCtx int) (string, error) {
	if backend == "ollama" && numCtx > 0 {
		return callOllamaWithCtx(model, system, user, numCtx)
	}
	return llmCall(backend, model, system, user)
}

// callOllama calls the local Ollama server.
func callOllama(model, system, user string) (string, error) {
	body := map[string]any{
		"model": model,
		"messages": []map[string]string{
			{"role": "system", "content": system},
			{"role": "user", "content": user},
		},
		"stream": false,
		"options": map[string]any{
			"temperature": 0,
			"num_predict": 2048,
		},
	}
	return doOllamaRequest("http://localhost:11434/api/chat", body)
}

// callOllamaWithCtx calls the local Ollama server with a custom num_ctx.
func callOllamaWithCtx(model, system, user string, numCtx int) (string, error) {
	body := map[string]any{
		"model": model,
		"messages": []map[string]string{
			{"role": "system", "content": system},
			{"role": "user", "content": user},
		},
		"stream": false,
		"options": map[string]any{
			"temperature": 0,
			"num_predict": 2048,
			"num_ctx":     numCtx,
		},
	}
	return doOllamaRequest("http://localhost:11434/api/chat", body)
}

func doOllamaRequest(url string, body any) (string, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("marshal ollama request: %w", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("ollama request: %w", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read ollama response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ollama %d: %s", resp.StatusCode, string(data))
	}
	var result struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return "", fmt.Errorf("parse ollama response: %w", err)
	}
	return result.Message.Content, nil
}

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

// callGemini calls the Google Gemini generateContent endpoint.
func callGemini(model, system, user string) (string, error) {
	apiKey, err := loadAPIKey("gemini")
	if err != nil {
		return "", err
	}
	endpoint := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s", model, apiKey)

	// Gemini merges system + user into a single user turn.
	combined := system + "\n\n" + user
	body := map[string]any{
		"contents": []map[string]any{
			{
				"role": "user",
				"parts": []map[string]string{
					{"text": combined},
				},
			},
		},
		"generationConfig": map[string]any{
			"temperature":     0,
			"maxOutputTokens": 2048,
		},
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("marshal gemini request: %w", err)
	}
	resp, err := http.Post(endpoint, "application/json", bytes.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("gemini request: %w", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read gemini response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("gemini %d: %s", resp.StatusCode, string(data))
	}
	var result struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return "", fmt.Errorf("parse gemini response: %w", err)
	}
	if len(result.Candidates) == 0 || len(result.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("gemini: empty response")
	}
	return result.Candidates[0].Content.Parts[0].Text, nil
}

// stripMarkdownFences removes wrapping markdown code fences from LLM output.
func stripMarkdownFences(s string) string {
	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, "```") {
		return s
	}
	// Remove opening fence (``` or ```yaml etc.)
	idx := strings.Index(s, "\n")
	if idx < 0 {
		return s
	}
	s = s[idx+1:]

	// Remove closing fence
	if last := strings.LastIndex(s, "```"); last >= 0 {
		s = s[:last]
	}
	return strings.TrimSpace(s)
}
