//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestRenderPyproject — pyproject.toml / requirements.txt / 의존성 목록 생성 검증
package scaffold

import (
	"strings"
	"testing"
)

func TestRenderRequirements(t *testing.T) {
	out, err := RenderRequirements()
	if err != nil {
		t.Fatalf("RenderRequirements error: %v", err)
	}
	for _, d := range runtimeDependencies() {
		if !strings.Contains(out, d) {
			t.Errorf("requirements missing %q\n%s", d, out)
		}
	}
	// dev deps should NOT be in requirements.txt
	if strings.Contains(out, "pytest>=") {
		t.Errorf("requirements.txt should not contain dev deps:\n%s", out)
	}
}
