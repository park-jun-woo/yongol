//ff:func feature=agent type=adapter control=sequence
//ff:what callOllama — 로컬 Ollama 서버에 chat completion 요청

package agent

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
