//ff:func feature=funcspec type=parser control=iteration dimension=1
//ff:what ParseDir 테스트 — 하위 디렉토리의 여러 func 스펙을 일괄 파싱

package funcspec

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseDir(t *testing.T) {
	dir := t.TempDir()

	authDir := filepath.Join(dir, "auth")
	os.MkdirAll(authDir, 0755)

	hash := `package auth

// @func hashPassword
// @description hash

type HashPasswordRequest struct {
	Password string
}
type HashPasswordResponse struct {
	HashedPassword string
}
func HashPassword(req HashPasswordRequest) (HashPasswordResponse, error) {
	return HashPasswordResponse{HashedPassword: "hashed"}, nil
}
`
	verify := `package auth

// @func verifyPassword
// @description verify

type VerifyPasswordRequest struct {
	PasswordHash string
	Password     string
}
type VerifyPasswordResponse struct{}
func VerifyPassword(req VerifyPasswordRequest) (VerifyPasswordResponse, error) {
	return VerifyPasswordResponse{}, nil
}
`
	os.WriteFile(filepath.Join(authDir, "hash_password.go"), []byte(hash), 0644)
	os.WriteFile(filepath.Join(authDir, "verify_password.go"), []byte(verify), 0644)

	specs, diags := ParseDir(dir)
	if len(diags) > 0 {
		t.Fatalf("ParseDir diagnostics: %v", diags)
	}
	if len(specs) != 2 {
		t.Fatalf("ParseDir count = %d, want 2", len(specs))
	}

	for _, s := range specs {
		if s.Package != "auth" {
			t.Errorf("Package = %q, want %q", s.Package, "auth")
		}
	}
}
