//ff:func feature=rule type=test control=sequence topic=catalog
//ff:what rulebookParseState feedLine/feedTableRow/appendDataRow/handleSectionHeading/resetTable 직접 검증
package catalog

import (
	"testing"
)

func TestFeedTableRow(t *testing.T) {
	st := &rulebookParseState{}
	// header first
	st.feedTableRow("| Rule ID | Level | Description | Source |")
	if !st.sawHeader {
		t.Fatalf("header not recognised")
	}
	// separator second
	st.feedTableRow("|---|---|---|---|")
	if !st.inTable {
		t.Fatalf("separator not recognised")
	}
	// data row
	st.feedTableRow("| D-02 | WARNING | warn | `s.go` |")
	if len(st.rules) != 1 || st.rules[0].ID != "D-02" {
		t.Errorf("data row not appended: %+v", st.rules)
	}
}
