//ff:func feature=agent type=adapter control=sequence
//ff:what callOllamaWithCtx — 커스텀 num_ctx로 Ollama 서버 호출

package agent

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
