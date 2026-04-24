//ff:func feature=gen-gogin type=test-helper control=sequence
//ff:what assertRenderLogValueFileOutput — 생성 소스 import / 라인 검증

package sqlcpost

import (
	"testing"
)

// assertRenderLogValueFileOutput verifies that the generated log-value source
// imports exactly "log/slog" and emits the expected slog attr lines (including
// REDACTED sensitive fields). Extracted from the sibling test to respect
// filefunc F1 / A12 (single control per func).
func assertRenderLogValueFileOutput(t *testing.T, src string) {
	t.Helper()
	assertRenderLogValueFileImports(t, src)
	assertRenderLogValueFileLines(t, src)
}
