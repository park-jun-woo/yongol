//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestNestDependencies_ZeroCov — 런타임 의존성 맵
package scaffold

import (
	"strings"
	"testing"
)

func TestRenderTSConfig_ZeroCov(t *testing.T) {
	out, err := RenderTSConfig()
	if err != nil {
		t.Fatalf("RenderTSConfig error: %v", err)
	}
	for _, want := range []string{`"module": "commonjs"`, `"target": "ES2021"`, "experimentalDecorators"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderTSConfig missing %q", want)
		}
	}
}
