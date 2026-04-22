//ff:func feature=rule type=test control=iteration dimension=1 topic=catalog
//ff:what test: Parse — rulebook.md 섹션/테이블 파싱 + Deprecated skip + Catalog Lookup 검증
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

// TestLoadCached verifies Load returns the same Catalog instance across calls.
func TestLoadCached(t *testing.T) {
	c1, err := Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	c2, err := Load()
	if err != nil {
		t.Fatalf("Load second call: %v", err)
	}
	if c1 != c2 {
		t.Errorf("Load should return cached Catalog, got distinct pointers")
	}
	if c1.Len() == 0 {
		t.Errorf("Load returned empty catalog")
	}
}

// TestParseInlineTable covers a compact, hand-authored rulebook snippet so
// that the parser behaviour stays focused regardless of the canonical
// rulebook's evolution.
func TestParseInlineTable(t *testing.T) {
	src := `# Rulebook
## A. Sample
prose

| Rule ID | Level | Description | Source |
|---|---|---|---|
| X-1 | ERROR | first | ` + "`pkg/x/x_01.go`" + ` |
| X-2 | WARNING | second | ` + "`pkg/x/x_02.go`" + ` |

## Deprecated

| Rule ID | Level | Description | Source |
|---|---|---|---|
| DEAD-1 | ERROR | gone | ` + "`pkg/x/dead.go`" + ` |
`
	rules, err := Parse(strings.NewReader(src))
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if len(rules) != 2 {
		t.Fatalf("expected 2 active rules, got %d: %+v", len(rules), rules)
	}
	if rules[0].ID != "X-1" || rules[0].Level != "ERROR" {
		t.Errorf("rules[0]: %+v", rules[0])
	}
	if rules[1].ID != "X-2" || rules[1].Level != "WARNING" {
		t.Errorf("rules[1]: %+v", rules[1])
	}
	if rules[0].Source != "pkg/x/x_01.go" {
		t.Errorf("rules[0].Source backtick strip: got %q", rules[0].Source)
	}
	if rules[0].SectionAnchor != "a-sample" {
		t.Errorf("section anchor: got %q, want a-sample", rules[0].SectionAnchor)
	}
}
