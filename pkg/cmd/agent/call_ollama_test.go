//ff:func feature=agent type=test control=sequence
//ff:what TestCallOllama — 로컬 Ollama 서버 부재 시 요청 에러 반환 검증

package agent

import "testing"

func TestCallOllama(t *testing.T) {
	// No Ollama server is expected at localhost:11434 in CI, so the request
	// fails fast (connection refused) and the error is propagated.
	out, err := callOllama("llama3", "system", "user")
	if err == nil {
		t.Skip("an Ollama server appears to be running locally; skipping negative-path assertion")
	}
	if out != "" {
		t.Errorf("expected empty output on error, got %q", out)
	}
}
