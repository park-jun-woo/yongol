//ff:func feature=rule type=test control=iteration dimension=1 topic=catalog
//ff:what TestParseCanonicalRulebook — embed rulebook.md 파싱 후 100+ 규칙 + 섹션 spot-check
package catalog

import (
	"bytes"
	"strings"
	"testing"
)

// TestParseCanonicalRulebook verifies the embedded rulebook.md parses cleanly
// and yields a non-trivial, deterministic catalog.
func TestParseCanonicalRulebook(t *testing.T) {
	rules, err := Parse(bytes.NewReader(Source()))
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if len(rules) < 100 {
		t.Fatalf("expected >=100 rules from canonical rulebook, got %d", len(rules))
	}

	// Spot-checks that protect against regressions.
	cat := NewCatalog(rules)
	spot := map[string]string{
		"S-27":   "A. SSaC Internal",
		"C-2":    "B. Manifest",
		"Q-1":    "D. Query / sqlc",
		"PRV-01": "S. Preserve",
	}
	for id, wantSection := range spot {
		meta, ok := cat.Lookup(id)
		if !ok {
			t.Errorf("rule %q not found in catalog", id)
			continue
		}
		if !strings.HasPrefix(meta.SectionTitle, wantSection) {
			t.Errorf("rule %q section: got %q, want prefix %q",
				id, meta.SectionTitle, wantSection)
		}
		if meta.Level != "ERROR" && meta.Level != "WARNING" {
			t.Errorf("rule %q level: got %q, want ERROR|WARNING", id, meta.Level)
		}
		if meta.Description == "" {
			t.Errorf("rule %q description is empty", id)
		}
		if meta.Source == "" {
			t.Errorf("rule %q source is empty", id)
		}
	}
}
