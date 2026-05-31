//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParsePolicyIfPresent — Policy 미탐지(return) + 탐지 시 ParsedPolicies 설정
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParsePolicyIfPresent_Present(t *testing.T) {
	dir := t.TempDir()
	content := "package authz\n\ndefault allow := false\n\nallow if {\n    input.action == \"Create\"\n}\n"
	if err := os.WriteFile(filepath.Join(dir, "authz.rego"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	fs := &Fullstack{}
	has := map[SSOTKind]DetectedSSOT{
		KindPolicy: {Kind: KindPolicy, Path: dir, Presence: SSOTPopulated},
	}
	parsePolicyIfPresent(fs, has)
	if len(fs.ParsedPolicies) == 0 {
		t.Fatalf("expected ParsedPolicies populated, diags=%+v", fs.ParseDiagnostics)
	}
}
