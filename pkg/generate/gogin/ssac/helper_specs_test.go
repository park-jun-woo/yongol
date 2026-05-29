//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what helperSpecs 단위 테스트 (포인터/deref 헬퍼 목록의 안정성·완전성)

package ssac

import (
	"strings"
	"testing"
)

func TestHelperSpecs(t *testing.T) {
	specs := helperSpecs()
	if len(specs) == 0 {
		t.Fatal("expected non-empty helper spec list")
	}

	// File names must be unique and every spec must carry file/what/code.
	seen := map[string]bool{}
	for _, h := range specs {
		if h.file == "" || h.what == "" || h.code == "" {
			t.Errorf("incomplete spec: %+v", h)
		}
		if !strings.HasSuffix(h.file, ".go") {
			t.Errorf("file %q should end in .go", h.file)
		}
		if seen[h.file] {
			t.Errorf("duplicate file name %q", h.file)
		}
		seen[h.file] = true
	}

	// Required helpers must be present.
	for _, want := range []string{"ptr_of.go", "deref_int.go", "deref_str.go", "deref_enum.go"} {
		if !seen[want] {
			t.Errorf("missing helper %q", want)
		}
	}

	// Stable order: regenerating yields identical first entry.
	if specs[0].file != helperSpecs()[0].file {
		t.Errorf("order is not stable")
	}
}
