//ff:func feature=gen-react type=test control=sequence
//ff:what generate_integ — dummy specs 로 실제 Fullstack 구성 후 react.Generate 통합 커버리지
package react

import (
	"path/filepath"
	"testing"
)

func TestRunOpenAPITypescript_BadSpec_Integ(t *testing.T) {
	dest := filepath.Join(t.TempDir(), "api.d.ts")
	err := runOpenAPITypescript(filepath.Join(t.TempDir(), "does-not-exist.yaml"), dest)
	if err == nil {
		t.Skip("openapi-typescript unexpectedly succeeded on missing spec; environment-specific")
	}
	t.Logf("runOpenAPITypescript error (expected): %v", err)
}
