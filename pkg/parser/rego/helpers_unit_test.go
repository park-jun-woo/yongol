//ff:func feature=policy type=test control=sequence
//ff:what TestRegoHelpers — unit tests for the pure rego parser helper functions
package rego

import (
	"reflect"
	"testing"
)

func TestLineOfOffset(t *testing.T) {
	s := "line1\nline2\nline3"
	tests := []struct {
		off  int
		want int
	}{
		{0, 1},
		{5, 1},          // index of the first '\n' is on line 1
		{6, 2},          // first char of line2
		{12, 3},         // first char of line3
		{len(s), 3},     // end-of-string
		{-1, 0},         // out of range low
		{len(s) + 1, 0}, // out of range high
	}
	for _, tt := range tests {
		if got := lineOfOffset(s, tt.off); got != tt.want {
			t.Errorf("lineOfOffset(off=%d) = %d, want %d", tt.off, got, tt.want)
		}
	}
}

func TestExtractClaimsRefs(t *testing.T) {
	content := `
allow if {
	input.claims.org_id == 5
	input.claims.role == "admin"
	input.claims.org_id == 6
}`
	p := &Policy{}
	extractClaimsRefs(content, p)
	// org_id appears twice but should be deduplicated; order is first-seen.
	want := []string{"org_id", "role"}
	if !reflect.DeepEqual(p.ClaimsRefs, want) {
		t.Errorf("ClaimsRefs = %v, want %v", p.ClaimsRefs, want)
	}
}

func TestExtractAllowRules(t *testing.T) {
	content := "package authz\n" +
		"allow {\n" +
		"\tinput.action == \"read\"\n" +
		"\tinput.resource == \"project\"\n" +
		"}\n" +
		"allow if {\n" +
		"\tinput.action == \"delete\"\n" +
		"\tinput.resource == \"project\"\n" +
		"\tinput.resource_owner == input.user.id\n" +
		"}\n"
	p := &Policy{File: "p.rego"}
	extractAllowRules(content, p)
	if len(p.Rules) != 2 {
		t.Fatalf("got %d rules, want 2: %+v", len(p.Rules), p.Rules)
	}
	if p.Rules[0].Actions[0] != "read" || p.Rules[0].Resource != "project" {
		t.Errorf("rule0 = %+v", p.Rules[0])
	}
	if p.Rules[1].Actions[0] != "delete" || !p.Rules[1].UsesOwner {
		t.Errorf("rule1 = %+v", p.Rules[1])
	}
	// SourceLine should be 1-based and increasing.
	if p.Rules[0].SourceLine <= 0 || p.Rules[1].SourceLine <= p.Rules[0].SourceLine {
		t.Errorf("source lines = %d, %d", p.Rules[0].SourceLine, p.Rules[1].SourceLine)
	}
}

func TestExtractErrorLine(t *testing.T) {
	// Non-OPA error → 0.
	if got := extractErrorLine(errNotOPA{}); got != 0 {
		t.Errorf("non-OPA error → %d, want 0", got)
	}
	// nil error → 0 (type assertion fails).
	if got := extractErrorLine(nil); got != 0 {
		t.Errorf("nil error → %d, want 0", got)
	}
}

type errNotOPA struct{}

func (errNotOPA) Error() string { return "plain error" }
