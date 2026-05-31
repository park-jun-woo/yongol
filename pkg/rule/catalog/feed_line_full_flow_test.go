//ff:func feature=rule type=test control=iteration dimension=1 topic=catalog
//ff:what rulebookParseState feedLine/feedTableRow/appendDataRow/handleSectionHeading/resetTable 직접 검증
package catalog

import (
	"testing"
)

func TestFeedLine_FullFlow(t *testing.T) {
	st := &rulebookParseState{}
	lines := []string{
		"## DDL Rules",
		"",
		"| Rule ID | Level | Description | Source |",
		"|---|---|---|---|",
		"| D-01 | ERROR | nullable | `a.go` |",
		"| D-02 | WARNING | serial | `b.go` |",
		"### Subsection",
		"## Deprecated",
		"| OLD-01 | ERROR | gone | `c.go` |",
	}
	for _, l := range lines {
		st.feedLine(l)
	}
	if len(st.rules) != 2 {
		t.Fatalf("expected 2 rules (deprecated skipped), got %d: %+v", len(st.rules), st.rules)
	}
	if st.rules[0].ID != "D-01" || st.rules[1].ID != "D-02" {
		t.Errorf("rules = %+v", st.rules)
	}
}
