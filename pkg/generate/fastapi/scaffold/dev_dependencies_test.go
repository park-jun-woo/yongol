//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestRenderPyproject — pyproject.toml / requirements.txt / 의존성 목록 생성 검증
package scaffold

import (
	"strings"
	"testing"
)

func TestDevDependencies(t *testing.T) {
	deps := devDependencies()
	if len(deps) == 0 {
		t.Fatal("expected dev dependencies")
	}
	found := false
	for _, d := range deps {
		if strings.HasPrefix(d, "pytest>=") {
			found = true
		}
	}
	if !found {
		t.Errorf("pytest dev dependency missing: %v", deps)
	}
}
