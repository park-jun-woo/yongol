//ff:func feature=manifest type=test control=sequence
//ff:what TestParseSentinelInserts_SemicolonInLiteral — 단일 인용 문자열 안의 `;` 는 terminator 아님

package ddl

import (
	"strings"
	"testing"
)

// TestParseSentinelInserts_SemicolonInLiteral pins the quote-aware
// terminator logic: a `;` inside a single-quoted literal must not end
// the statement early.
func TestParseSentinelInserts_SemicolonInLiteral(t *testing.T) {
	content := `-- @sentinel
INSERT INTO t (id, msg) VALUES (0, 'hi; there') ON CONFLICT DO NOTHING;
`
	results := parseSentinelInserts(content)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if !strings.Contains(results[0].SQL, "hi; there") {
		t.Errorf("literal semicolon was lost: %q", results[0].SQL)
	}
	if !strings.HasSuffix(strings.TrimSpace(results[0].SQL), ";") {
		t.Errorf("terminator ; missing: %q", results[0].SQL)
	}
}
