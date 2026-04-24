//ff:func feature=gen-gogin type=test-helper control=iteration dimension=1
//ff:what assertRenderLogValueFileLines — slog.Any / REDACTED 라인 검증

package sqlcpost

import (
	"strings"
	"testing"
)

// assertRenderLogValueFileLines verifies that the LogValue body contains the
// expected slog.Any / slog.String lines in the generated source.
func assertRenderLogValueFileLines(t *testing.T, src string) {
	t.Helper()
	for _, line := range []string{
		"\t\tslog.Any(\"id\", r.ID),\n",
		"\t\tslog.Any(\"org_id\", r.OrgID),\n",
		"\t\tslog.Any(\"email\", r.Email),\n",
		"\t\tslog.String(\"password_hash\", \"[REDACTED]\"),\n",
		"\t\tslog.Any(\"role\", r.Role),\n",
		"\t\tslog.Any(\"created_at\", r.CreatedAt),\n",
	} {
		if !strings.Contains(src, line) {
			t.Errorf("missing line %q in generated source:\n%s", line, src)
		}
	}
}
