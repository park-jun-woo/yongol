//ff:func feature=cli type=helper control=selection
//ff:what parseModelFlag — "backend:model" 문자열 파싱 (ollama/xai/gemini 지원)
package main

import (
	"fmt"
	"strings"
)

// parseModelFlag splits "backend:model" into backend and model name.
// For ollama models with colons in the name (e.g. "ollama:gemma4:e4b"),
// the backend is the first segment and the model is everything after.
func parseModelFlag(flag string) (backend, model string, err error) {
	idx := strings.Index(flag, ":")
	if idx < 0 {
		return "", "", fmt.Errorf("invalid --model %q: expected format backend:model (e.g. ollama:gemma4:e4b)", flag)
	}
	backend = flag[:idx]
	model = flag[idx+1:]
	switch backend {
	case "ollama", "xai", "gemini":
		// ok
	default:
		return "", "", fmt.Errorf("invalid --model backend %q: supported backends: ollama, xai, gemini", backend)
	}
	if model == "" {
		return "", "", fmt.Errorf("invalid --model %q: model name is empty", flag)
	}
	return backend, model, nil
}
