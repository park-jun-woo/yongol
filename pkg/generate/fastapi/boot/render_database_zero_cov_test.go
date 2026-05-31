//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestHasActiveBlock_ZeroCov — 활성 블록 존재 여부
package boot

import (
	"strings"
	"testing"
)

func TestRenderDatabase_ZeroCov(t *testing.T) {
	out, err := RenderDatabase()
	if err != nil {
		t.Fatalf("RenderDatabase error: %v", err)
	}
	for _, want := range []string{"create_async_engine", "async_session", "class Base(DeclarativeBase)"} {
		if !strings.Contains(out, want) {
			t.Errorf("RenderDatabase missing %q", want)
		}
	}
}
