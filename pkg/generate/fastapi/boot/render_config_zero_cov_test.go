//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestHasActiveBlock_ZeroCov — 활성 블록 존재 여부
package boot

import (
	"strings"
	"testing"
)

func TestRenderConfig_ZeroCov(t *testing.T) {
	out, err := RenderConfig()
	if err != nil {
		t.Fatalf("RenderConfig error: %v", err)
	}
	for _, want := range []string{"BaseSettings", "database_url", "jwt_secret", "settings = Settings()"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderConfig missing %q", want)
		}
	}
}
