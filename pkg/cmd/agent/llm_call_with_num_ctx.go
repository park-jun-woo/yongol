//ff:func feature=agent type=adapter control=sequence
//ff:what llmCallWithNumCtx — ollama num_ctx 지정 LLM 호출

package agent

// llmCallWithNumCtx calls the LLM with a custom num_ctx for ollama.
// For non-ollama backends, numCtx is ignored and the standard call is used.
func llmCallWithNumCtx(backend, model, system, user string, numCtx int) (string, error) {
	if backend == "ollama" && numCtx > 0 {
		return callOllamaWithCtx(model, system, user, numCtx)
	}
	return llmCall(backend, model, system, user)
}
