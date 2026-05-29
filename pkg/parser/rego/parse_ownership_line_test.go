//ff:func feature=policy type=parser control=sequence
//ff:what OwnershipMapping.SourceLine 이 @ownership 어노테이션 줄 번호로 채워지는지 검증

package rego

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParsePolicyFile_OwnershipSourceLine(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "gig.rego")

	// Lines (1-based):
	// 1: package authz.gig
	// 2: (empty)
	// 3: # @ownership gig: gigs.client_id
	// 4: (empty)
	// 5: # @ownership proposal: proposals.freelancer_id via gigs.client_id
	// 6: (empty)
	// 7: default allow := false
	content := "package authz.gig\n" +
		"\n" +
		"# @ownership gig: gigs.client_id\n" +
		"\n" +
		"# @ownership proposal: proposals.freelancer_id via gigs.client_id\n" +
		"\n" +
		"default allow := false\n"

	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	p, _ := ParsePolicyFile(path)
	if p == nil {
		t.Fatal("ParsePolicyFile returned nil policy")
	}
	if len(p.Ownerships) != 2 {
		t.Fatalf("Ownerships count = %d, want 2", len(p.Ownerships))
	}

	if p.Ownerships[0].Resource != "gig" {
		t.Errorf("Ownerships[0].Resource = %q, want %q", p.Ownerships[0].Resource, "gig")
	}
	if p.Ownerships[0].SourceLine != 3 {
		t.Errorf("Ownerships[0].SourceLine = %d, want 3", p.Ownerships[0].SourceLine)
	}

	if p.Ownerships[1].Resource != "proposal" {
		t.Errorf("Ownerships[1].Resource = %q, want %q", p.Ownerships[1].Resource, "proposal")
	}
	if p.Ownerships[1].SourceLine != 5 {
		t.Errorf("Ownerships[1].SourceLine = %d, want 5", p.Ownerships[1].SourceLine)
	}
	if p.Ownerships[1].JoinTable != "gigs" || p.Ownerships[1].JoinFK != "client_id" {
		t.Errorf("Ownerships[1] join = %q.%q, want gigs.client_id", p.Ownerships[1].JoinTable, p.Ownerships[1].JoinFK)
	}
}
