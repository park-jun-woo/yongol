//ff:func feature=agent type=test control=sequence
//ff:what TestCallOllamaWithCtx — 커스텀 num_ctx 호출 시 서버 부재 에러 반환 검증

package agent

import "testing"

func TestCallOllamaWithCtx(t *testing.T) {
	// No Ollama server at localhost:11434 → request fails fast and is propagated.
	out, err := callOllamaWithCtx("llama3", "system", "user", 8192)
	if err == nil {
		t.Skip("an Ollama server appears to be running locally; skipping negative-path assertion")
	}
	if out != "" {
		t.Errorf("expected empty output on error, got %q", out)
	}
}
