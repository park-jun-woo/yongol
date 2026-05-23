//ff:func feature=agent type=adapter control=selection
//ff:what llmCall — LLM 호출 추상화 (ollama/xai/gemini)

package agent

import "fmt"

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
