//ff:func feature=policy type=test control=sequence
//ff:what Rego helper (parseActionSet / findClosingBrace / looksLikeOwnership / parseOwnershipLine / processAllowBlock) 단위 회귀

package rego

import "testing"

func TestParseActionSet_Basic(t *testing.T) {
	got := parseActionSet(` "Create" , "Update" , "Delete" `)
	if len(got) != 3 {
		t.Fatalf("got %v", got)
	}
	if got[0] != "Create" || got[2] != "Delete" {
		t.Errorf("got %v", got)
	}
}

func TestParseActionSet_Empty(t *testing.T) {
	if got := parseActionSet(""); got != nil {
		t.Errorf("got %v, want nil", got)
	}
}

func TestFindClosingBrace_Simple(t *testing.T) {
	s := "abc}rest"
	if got := findClosingBrace(s); got != 3 {
		t.Errorf("got %d, want 3", got)
	}
}

func TestFindClosingBrace_Nested(t *testing.T) {
	s := "{inner}more}tail"
	if got := findClosingBrace(s); got != 11 {
		t.Errorf("got %d, want 11", got)
	}
}

func TestFindClosingBrace_Unbalanced(t *testing.T) {
	s := "no closing"
	if got := findClosingBrace(s); got != -1 {
		t.Errorf("got %d, want -1", got)
	}
}

func TestLooksLikeOwnership(t *testing.T) {
	cases := map[string]bool{
		"# @ownership gig: gigs.client_id": true,
		"#@ownership":                      true,
		"# @ownership":                     true,
		"# normal comment":                 false,
		"allow {":                          false,
		"":                                 false,
	}
	for in, want := range cases {
		if got := looksLikeOwnership(in); got != want {
			t.Errorf("looksLikeOwnership(%q) = %v, want %v", in, got, want)
		}
	}
}

func TestParseOwnershipLine_WithVia(t *testing.T) {
	om, ok := parseOwnershipLine("# @ownership proposal: proposals.freelancer_id via gigs.client_id")
	if !ok {
		t.Fatal("expected ok")
	}
	if om.Resource != "proposal" || om.Table != "proposals" || om.Column != "freelancer_id" {
		t.Errorf("base = %+v", om)
	}
	if om.JoinTable != "gigs" || om.JoinFK != "client_id" {
		t.Errorf("join = %+v", om)
	}
}

func TestParseOwnershipLine_Simple(t *testing.T) {
	om, ok := parseOwnershipLine("# @ownership gig: gigs.client_id")
	if !ok {
		t.Fatal("expected ok")
	}
	if om.JoinTable != "" || om.JoinFK != "" {
		t.Errorf("join should be empty: %+v", om)
	}
}

func TestParseOwnershipLine_Invalid(t *testing.T) {
	if _, ok := parseOwnershipLine("# not an ownership"); ok {
		t.Error("expected false")
	}
}

func TestProcessAllowBlock_Basic(t *testing.T) {
	block := `
    input.action == "Read"
    input.resource == "note"
`
	rule, ok := processAllowBlock(block)
	if !ok {
		t.Fatal("expected ok")
	}
	if len(rule.Actions) != 1 || rule.Actions[0] != "Read" {
		t.Errorf("Actions = %v", rule.Actions)
	}
	if rule.Resource != "note" {
		t.Errorf("Resource = %q", rule.Resource)
	}
}

func TestProcessAllowBlock_MissingResource(t *testing.T) {
	block := `    input.action == "Read"`
	_, ok := processAllowBlock(block)
	if ok {
		t.Errorf("expected false when resource missing")
	}
}
