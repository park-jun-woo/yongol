//ff:func feature=rule type=test control=sequence topic=catalog
//ff:what TestCatalogHelpers — unit tests for the pure rule catalog helper functions
package catalog

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplitRow(t *testing.T) {
	got := splitRow("| a | b | c |")
	want := []string{" a ", " b ", " c "}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("splitRow = %q, want %q", got, want)
	}
	// Without surrounding pipes.
	if got := splitRow("x|y"); !reflect.DeepEqual(got, []string{"x", "y"}) {
		t.Errorf("no-pipe-edge = %q", got)
	}
}

func TestIsRuleTableHeader(t *testing.T) {
	if !isRuleTableHeader("| Rule ID | Level | Description | Source |") {
		t.Error("expected canonical header recognised")
	}
	// Case-insensitive.
	if !isRuleTableHeader("| rule id | LEVEL | description | source |") {
		t.Error("expected case-insensitive match")
	}
	// Wrong columns.
	if isRuleTableHeader("| ID | Name |") {
		t.Error("non-header should be rejected")
	}
	// Too few cells.
	if isRuleTableHeader("| Rule ID | Level |") {
		t.Error("too few cells should be rejected")
	}
}

func TestIsTableSeparator(t *testing.T) {
	if !isTableSeparator("|---|---|---|---|") {
		t.Error("expected separator recognised")
	}
	if !isTableSeparator("| :--- | ---: | :---: |") {
		t.Error("alignment colons should be accepted")
	}
	// Empty cell → not a separator.
	if isTableSeparator("|---||---|") {
		t.Error("empty cell should reject")
	}
	// Non-dash content → not a separator.
	if isTableSeparator("| abc | def |") {
		t.Error("text content should reject")
	}
}

func TestSectionAnchor(t *testing.T) {
	tests := []struct{ in, want string }{
		{"D-Series Rules", "d-series-rules"},
		{"XQS.20 Return Type", "xqs-20-return-type"},
		{"Trailing  ", "trailing"},
		{"a/b c", "a-b-c"},
		{"!!!", ""}, // all punctuation stripped
	}
	for _, tt := range tests {
		if got := sectionAnchor(tt.in); got != tt.want {
			t.Errorf("sectionAnchor(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestWriteSectionAnchorRune(t *testing.T) {
	// Alphanumeric is written, resets prevDash.
	var b strings.Builder
	if dash := writeSectionAnchorRune(&b, 'a', false); dash {
		t.Error("alnum should return prevDash=false")
	}
	if b.String() != "a" {
		t.Errorf("buffer = %q", b.String())
	}
	// Space after content writes a dash.
	if dash := writeSectionAnchorRune(&b, ' ', false); !dash {
		t.Error("space after content should write dash")
	}
	if b.String() != "a-" {
		t.Errorf("buffer = %q, want a-", b.String())
	}
	// Leading separator (empty buffer) is suppressed.
	var b2 strings.Builder
	if dash := writeSectionAnchorRune(&b2, '-', false); dash {
		t.Error("leading separator should not write dash")
	}
	if b2.Len() != 0 {
		t.Errorf("leading separator wrote %q", b2.String())
	}
	// Unknown punctuation is dropped, prevDash preserved.
	var b3 strings.Builder
	b3.WriteString("x")
	if dash := writeSectionAnchorRune(&b3, '!', false); dash {
		t.Error("punctuation should preserve prevDash=false")
	}
	if b3.String() != "x" {
		t.Errorf("punctuation altered buffer: %q", b3.String())
	}
}
