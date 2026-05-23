//ff:func feature=agent type=adapter control=sequence
//ff:what doOllamaRequest — Ollama HTTP 요청 전송 및 응답 파싱

package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
