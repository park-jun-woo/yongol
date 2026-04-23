//ff:func feature=rule type=test control=iteration dimension=1 topic=catalog
//ff:what TestParseSkipsDeprecatedSection — `## Deprecated` 섹션의 규칙이 카탈로그에서 제외됨
package catalog

import (
	"bytes"
	"testing"
)

// TestParseSkipsDeprecatedSection ensures rules listed under `## Deprecated`
// are intentionally excluded from the active catalog.
func TestParseSkipsDeprecatedSection(t *testing.T) {
	rules, err := Parse(bytes.NewReader(Source()))
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	cat := NewCatalog(rules)
	// XDO-1 and S-52 are explicitly listed under the Deprecated section in
	// the canonical rulebook. They must not leak into the active catalog.
	for _, id := range []string{"XDO-1", "S-52", "S-55", "M-1"} {
		if _, ok := cat.Lookup(id); ok {
			t.Errorf("deprecated rule %q should be skipped", id)
		}
	}
}
