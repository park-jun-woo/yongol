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
		"INI-01":  "INI. Init Check",
		"S-27":    "A. SSaC Internal",
		"C-2":     "B. Manifest",
		"SEC-404": "B. Manifest",
		"Q-01":    "D. Query / sqlc",
		"V-01":    "Z1. Design Internal",
		"PRV-01":  "S. Preserve",
		"TM-17":   "U. STML",
		"TM-26":   "U. STML",
		"FT-01":   "V. Features Internal",
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

	// TM-17 regression: the rulebook row escapes its pipes as `\|\|` (GFM);
	// splitRow must restore the literal `||` code span instead of splitting
	// the description cell and blanking the source.
	if meta, ok := cat.Lookup("TM-17"); ok {
		if !strings.Contains(meta.Description, "`||`") {
			t.Errorf("TM-17 description lost its `||` code span: %q", meta.Description)
		}
		if meta.Source != "pkg/validate/stml_openapi/tm_17_guard_syntax.go" {
			t.Errorf("TM-17 source = %q, want tm_17_guard_syntax.go path", meta.Source)
		}
	}
}
