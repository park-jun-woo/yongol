//ff:func feature=rule type=test control=sequence topic=catalog
//ff:what rulebookParseState feedLine/feedTableRow/appendDataRow/handleSectionHeading/resetTable 직접 검증
package catalog

import (
	"testing"
)

func TestAppendDataRow(t *testing.T) {
	st := &rulebookParseState{sectionTitle: "DDL Rules"}
	st.appendDataRow("| D-01 | ERROR | must not be null | `pkg/x.go` |")
	if len(st.rules) != 1 {
		t.Fatalf("rules len = %d, want 1", len(st.rules))
	}
	r := st.rules[0]
	if r.ID != "D-01" || r.Level != "ERROR" || r.Description != "must not be null" || r.Source != "pkg/x.go" {
		t.Errorf("rule = %+v", r)
	}
	if r.SectionTitle != "DDL Rules" || r.SectionAnchor != "ddl-rules" {
		t.Errorf("section = %q / %q", r.SectionTitle, r.SectionAnchor)
	}

	// invalid rows are ignored: too few cells, empty id, bad level
	st.appendDataRow("| only | two |")
	st.appendDataRow("|  | ERROR | d | s |")
	st.appendDataRow("| X | INFO | d | s |")
	if len(st.rules) != 1 {
		t.Errorf("invalid rows should be ignored, got %d", len(st.rules))
	}
}
