//ff:func feature=funcspec type=parser control=sequence
//ff:what ParseFile 단일 import 파싱 테스트 — Imports 필드에 경로가 수집되는지 검증

package funcspec

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFileImports(t *testing.T) {
	dir := t.TempDir()

	src := `package auth

import "golang.org/x/crypto/bcrypt"

// @func hashPassword
// @description 평문 비밀번호를 bcrypt 해시로 변환한다

type HashPasswordRequest struct {
	Password string
}

type HashPasswordResponse struct {
	HashedPassword string
}

func HashPassword(req HashPasswordRequest) (HashPasswordResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	return HashPasswordResponse{HashedPassword: string(hash)}, err
}
`
	path := filepath.Join(dir, "hash_password.go")
	os.WriteFile(path, []byte(src), 0644)

	spec, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("ParseFile diagnostics: %v", diags)
	}
	if spec == nil {
		t.Fatal("expected non-nil spec")
	}

	if len(spec.Imports) != 1 {
		t.Fatalf("Imports count = %d, want 1", len(spec.Imports))
	}
	if spec.Imports[0] != "golang.org/x/crypto/bcrypt" {
		t.Errorf("Imports[0] = %q, want %q", spec.Imports[0], "golang.org/x/crypto/bcrypt")
	}
}
