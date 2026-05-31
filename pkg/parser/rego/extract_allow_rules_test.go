//ff:func feature=policy type=test control=sequence
//ff:what TestRegoHelpers — unit tests for the pure rego parser helper functions
package rego

import (
	"testing"
)

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
