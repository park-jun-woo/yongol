//ff:func feature=rule type=test control=iteration dimension=1 topic=catalog
//ff:what rulebookParseState feedLine/feedTableRow/appendDataRow/handleSectionHeading/resetTable 직접 검증

package catalog

import "testing"

func TestResetTable(t *testing.T) {
	st := &rulebookParseState{inTable: true, sawHeader: true}
	st.resetTable()
	if st.inTable || st.sawHeader {
		t.Errorf("resetTable left flags set: %+v", st)
	}
}

func TestHandleSectionHeading(t *testing.T) {
	st := &rulebookParseState{inTable: true, sawHeader: true}
	st.handleSectionHeading("## SSaC Rules")
	if st.sectionTitle != "SSaC Rules" || st.sectionSkip {
		t.Errorf("section state = %+v", st)
	}
	if st.inTable || st.sawHeader {
		t.Errorf("flags should reset on heading: %+v", st)
	}
	// Deprecated section sets skip
	st.handleSectionHeading("## Deprecated")
	if !st.sectionSkip {
		t.Errorf("Deprecated should set sectionSkip")
	}
}

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

func TestFeedTableRow_UnexpectedBetweenHeaderAndSeparator(t *testing.T) {
	st := &rulebookParseState{sawHeader: true}
	st.feedTableRow("| not a separator | x |")
	if st.sawHeader {
		t.Errorf("should bail (reset sawHeader) on unexpected row")
	}
}

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

func TestFeedLine_NonTableResetsTable(t *testing.T) {
	st := &rulebookParseState{inTable: true, sawHeader: true}
	st.feedLine("just a paragraph")
	if st.inTable || st.sawHeader {
		t.Errorf("non-table line should reset table flags")
	}
}
