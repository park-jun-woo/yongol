//ff:func feature=gen-gogin type=test-helper control=iteration dimension=1
//ff:what assertRenderLogValueFileImports — import 블록 엄격 검증

package sqlcpost

import (
	"strings"
	"testing"
)

// assertRenderLogValueFileImports checks that only "log/slog" is imported
// and no time / json references leak into the generated source.
func assertRenderLogValueFileImports(t *testing.T, src string) {
	t.Helper()
	wantImportBlock := "import (\n\t\"log/slog\"\n)\n\n"
	if !strings.Contains(src, wantImportBlock) {
		t.Errorf("expected import block %q in output, got:\n%s", wantImportBlock, src)
	}
	for _, banned := range []string{"\"time\"", "\"encoding/json\"", "time.Time", "json.RawMessage"} {
		if strings.Contains(src, banned) {
			t.Errorf("generated source must not reference %s anymore (BUG-024):\n%s", banned, src)
		}
	}
}
