//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestRenderPyproject — pyproject.toml / requirements.txt / 의존성 목록 생성 검증
package scaffold

import (
	"strings"
	"testing"
)

func TestRuntimeDependencies(t *testing.T) {
	deps := runtimeDependencies()
	if len(deps) == 0 {
		t.Fatal("expected runtime dependencies")
	}
	found := false
	for _, d := range deps {
		if strings.HasPrefix(d, "fastapi>=") {
			found = true
		}
	}
	if !found {
		t.Errorf("fastapi dependency missing: %v", deps)
	}
}
