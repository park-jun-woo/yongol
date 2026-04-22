//ff:func feature=policy type=test control=sequence
//ff:what extractAllowRules — 액션 집합 / role / owner 참조 검출 + AllowRule.SourceLine 회귀

package rego

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParsePolicyFile_ActionSetAndRole(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "authz.rego")

	// "allow if {" legacy form — extractAllowRules 가 "allow {" 도 커버한다.
	content := `package authz

default allow := false

allow if {
    input.action in {"Create","Update"}
    input.resource == "workflow"
    input.user.role == "editor"
}

allow if {
    input.action == "Delete"
    input.resource == "workflow"
    input.resource_owner == input.claims.sub
}
`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	p, diags := ParsePolicyFile(path)
	if p == nil {
		t.Fatalf("nil policy; diags=%v", diags)
	}
	if len(p.Rules) != 2 {
		t.Fatalf("Rules count = %d, want 2", len(p.Rules))
	}

	r0 := p.Rules[0]
	if len(r0.Actions) != 2 {
		t.Errorf("Rules[0].Actions = %v, want 2", r0.Actions)
	}
	if r0.Resource != "workflow" {
		t.Errorf("Rules[0].Resource = %q", r0.Resource)
	}
	if !r0.UsesRole || r0.RoleValue != "editor" {
		t.Errorf("Rules[0] role: UsesRole=%v Value=%q", r0.UsesRole, r0.RoleValue)
	}
	if r0.UsesOwner {
		t.Errorf("Rules[0].UsesOwner should be false")
	}

	r1 := p.Rules[1]
	if len(r1.Actions) != 1 || r1.Actions[0] != "Delete" {
		t.Errorf("Rules[1].Actions = %v", r1.Actions)
	}
	if !r1.UsesOwner {
		t.Errorf("Rules[1].UsesOwner = false, want true")
	}

	// claims.sub 참조 수집
	if len(p.ClaimsRefs) == 0 {
		t.Errorf("ClaimsRefs empty, want at least [sub]")
	}
	foundSub := false
	for _, c := range p.ClaimsRefs {
		if c == "sub" {
			foundSub = true
		}
	}
	if !foundSub {
		t.Errorf("ClaimsRefs %v missing 'sub'", p.ClaimsRefs)
	}
}

func TestParsePolicyFile_MalformedOwnership(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.rego")
	content := `package authz

# @ownership this is malformed

default allow := false
`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	_, diags := ParsePolicyFile(path)
	var gotMalformed bool
	for _, d := range diags {
		if d.Line == 3 {
			gotMalformed = true
		}
	}
	if !gotMalformed {
		t.Errorf("expected malformed @ownership diag at line 3, got %v", diags)
	}
}
