//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what serverHelperSources 단위 테스트 (헬퍼 spec → fileName→소스 맵)

package ssac

import (
	"strings"
	"testing"
)

func TestServerHelperSources(t *testing.T) {
	sources := serverHelperSources()
	specs := helperSpecs()
	if len(sources) != len(specs) {
		t.Fatalf("source count %d != spec count %d", len(sources), len(specs))
	}
	for _, h := range specs {
		body, ok := sources[h.file]
		if !ok {
			t.Errorf("missing source for %q", h.file)
			continue
		}
		if !strings.Contains(body, "package service") {
			t.Errorf("%q missing package clause", h.file)
		}
		if !strings.Contains(body, strings.TrimSpace(h.code)) {
			t.Errorf("%q missing function body", h.file)
		}
	}
}
