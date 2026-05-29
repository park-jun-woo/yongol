//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what test: TestGenerateUserClaim — internal/model/user_claim.go emit 결과 검증

package auth

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestGenerateUserClaim verifies the emitted model/user_claim.go file
// carries the expected package declaration, struct name, and JSON-tagged
// fields in sorted-by-name order. The fields mirror what parseClaims
// would produce from a realistic manifest.backend.auth.claims map.
func TestGenerateUserClaim(t *testing.T) {
	dir := t.TempDir()
	fields := []ClaimField{
		{Name: "Email", Key: "email", GoType: "string"},
		{Name: "ID", Key: "user_id", GoType: "int64"},
		{Name: "OrgID", Key: "org_id", GoType: "int64"},
		{Name: "Role", Key: "role", GoType: "string"},
	}
	if err := generateUserClaim(dir, fields); err != nil {
		t.Fatalf("generateUserClaim: %v", err)
	}
	path := filepath.Join(dir, "backend", "internal", "model", "user_claim.go")
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read emitted file: %v", err)
	}
	got := string(raw)

	wantSubstrings := []string{
		"package model",
		"type UserClaim struct {",
		"Email string `json:\"email\"`",
		"ID int64 `json:\"user_id\"`",
		"OrgID int64 `json:\"org_id\"`",
		"Role string `json:\"role\"`",
		"//ff:type feature=model type=model",
	}
	for _, want := range wantSubstrings {
		if !strings.Contains(got, want) {
			t.Errorf("emitted user_claim.go missing %q\n---\n%s", want, got)
		}
	}

	// Previous generators emitted Claim / CurrentUser; those names must
	// no longer appear anywhere in the UserClaim output.
	forbid := []string{"type Claim struct", "type CurrentUser struct"}
	for _, bad := range forbid {
		if strings.Contains(got, bad) {
			t.Errorf("emitted user_claim.go unexpectedly contains %q", bad)
		}
	}
}
