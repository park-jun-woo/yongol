//ff:func feature=validate type=util control=sequence topic=query-structural
//ff:what headTokenEquals — RawType 의 head (배열/파라미터/다중 단어 정규화 후) 와 want 를 case-insensitive 비교

package query

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// headTokenEquals upper-cases the head token of raw (stripping "[]" and
// "(...)"), normalises multi-word PostgreSQL type names to their
// canonical single-token alias via ddl.NormalizePGTypeHead
// ("DOUBLE PRECISION" → "FLOAT8", "TIMESTAMP WITH TIME ZONE" →
// "TIMESTAMPTZ" etc.), and compares case-insensitively to want. Shared
// by Q-12 ~ Q-18 column filters so a column declared in either form
// triggers the same per-type override rule.
func headTokenEquals(raw, want string) bool {
	t := strings.TrimSpace(raw)
	if strings.HasSuffix(t, "[]") {
		t = strings.TrimSpace(strings.TrimSuffix(t, "[]"))
	}
	if idx := strings.Index(t, "("); idx > 0 {
		t = strings.TrimSpace(t[:idx])
	}
	head := ddl.NormalizePGTypeHead(t)
	return strings.EqualFold(head, want)
}
