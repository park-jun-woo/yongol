//ff:func feature=policy type=test control=iteration dimension=1
//ff:what ParsePolicyFile — 잘못된 @ownership 주석은 3라인 진단 발행

package rego

import (
	"os"
	"path/filepath"
	"testing"
)

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
