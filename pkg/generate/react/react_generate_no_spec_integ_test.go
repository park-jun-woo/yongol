//ff:func feature=gen-react type=test control=sequence
//ff:what generate_integ — dummy specs 로 실제 Fullstack 구성 후 react.Generate 통합 커버리지
package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestReactGenerate_NoSpec_Integ(t *testing.T) {
	fs := &yongol.Fullstack{} // empty: SpecsDir "", OpenAPIDoc nil, Manifest nil
	out := t.TempDir()
	if err := Generate(fs, out); err != nil {
		t.Fatalf("Generate with empty fullstack: %v", err)
	}
	stub := filepath.Join(out, "frontend", "src", "types", "api.d.ts")
	data, err := os.ReadFile(stub)
	if err != nil {
		t.Fatalf("expected api.d.ts stub: %v", err)
	}
	if len(data) == 0 {
		t.Error("api.d.ts stub is empty")
	}
}
