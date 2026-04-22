//ff:func feature=policy type=parser control=sequence
//ff:what AllowRule.SourceLine 이 allow 규칙 시작 줄 번호로 채워지는지 검증

package rego

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParsePolicyFile_AllowRuleSourceLine(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "authz.rego")

	// Lines (1-based):
	// 1: package authz
	// 2: (empty)
	// 3: default allow := false
	// 4: (empty)
	// 5: allow if {
	// 6:     input.action == "Create"
	// 7:     input.resource == "workflow"
	// 8: }
	// 9: (empty)
	// 10: allow if {
	// 11:     input.action == "Delete"
	// 12:     input.resource == "workflow"
	// 13: }
	content := "package authz\n" +
		"\n" +
		"default allow := false\n" +
		"\n" +
		"allow if {\n" +
		"    input.action == \"Create\"\n" +
		"    input.resource == \"workflow\"\n" +
		"}\n" +
		"\n" +
		"allow if {\n" +
		"    input.action == \"Delete\"\n" +
		"    input.resource == \"workflow\"\n" +
		"}\n"

	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	p, diags := ParsePolicyFile(path)
	if p == nil {
		t.Fatalf("ParsePolicyFile returned nil policy; diags=%v", diags)
	}
	if len(p.Rules) != 2 {
		t.Fatalf("Rules count = %d, want 2", len(p.Rules))
	}

	if got, want := p.Rules[0].SourceLine, 5; got != want {
		t.Errorf("Rules[0].SourceLine = %d, want %d", got, want)
	}
	if got, want := p.Rules[1].SourceLine, 10; got != want {
		t.Errorf("Rules[1].SourceLine = %d, want %d", got, want)
	}
}
