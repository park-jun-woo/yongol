//ff:func feature=rule type=test control=sequence topic=catalog
//ff:what TestParseInlineTable — 손으로 작성한 rulebook 스니펫으로 Parse 거동 고정
package catalog

import (
	"strings"
	"testing"
)

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
