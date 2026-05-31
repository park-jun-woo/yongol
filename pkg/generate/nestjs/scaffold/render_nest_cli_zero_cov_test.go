//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestNestDependencies_ZeroCov — 런타임 의존성 맵
package scaffold

import (
	"strings"
	"testing"
)

func TestRenderNestCLI_ZeroCov(t *testing.T) {
	out, err := RenderNestCLI()
	if err != nil {
		t.Fatalf("RenderNestCLI error: %v", err)
	}
	for _, want := range []string{"@nestjs/schematics", "sourceRoot", "deleteOutDir"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderNestCLI missing %q", want)
		}
	}
}
