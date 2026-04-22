//ff:func feature=funcspec type=parser control=sequence
//ff:what ParseFile 기본 동작 테스트 — @func 어노테이션, Request/Response 필드, HasBody 검증

package funcspec

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFile(t *testing.T) {
	dir := t.TempDir()
	pkgDir := filepath.Join(dir, "auth")
	os.MkdirAll(pkgDir, 0755)

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
	path := filepath.Join(pkgDir, "hash_password.go")
	os.WriteFile(path, []byte(src), 0644)

	spec, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("ParseFile diagnostics: %v", diags)
	}
	if spec == nil {
		t.Fatal("expected non-nil spec")
	}

	if spec.Name != "hashPassword" {
		t.Errorf("Name = %q, want %q", spec.Name, "hashPassword")
	}
	if spec.Description != "평문 비밀번호를 bcrypt 해시로 변환한다" {
		t.Errorf("Description = %q", spec.Description)
	}
	if len(spec.RequestFields) != 1 {
		t.Fatalf("RequestFields count = %d, want 1", len(spec.RequestFields))
	}
	if spec.RequestFields[0].Name != "Password" || spec.RequestFields[0].Type != "string" {
		t.Errorf("RequestFields[0] = %+v", spec.RequestFields[0])
	}
	if len(spec.ResponseFields) != 1 {
		t.Fatalf("ResponseFields count = %d, want 1", len(spec.ResponseFields))
	}
	if spec.ResponseFields[0].Name != "HashedPassword" || spec.ResponseFields[0].Type != "string" {
		t.Errorf("ResponseFields[0] = %+v", spec.ResponseFields[0])
	}
	if !spec.HasBody {
		t.Error("HasBody = false, want true")
	}
}
