//ff:func feature=funcspec type=parser control=sequence
//ff:what ParseDir 분리 파일 테스트 — Request/Response struct가 별도 파일에 있을 때 패키지 병합 파싱

package funcspec

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseDirSplitFiles(t *testing.T) {
	dir := t.TempDir()
	authDir := filepath.Join(dir, "auth")
	os.MkdirAll(authDir, 0755)

	// @func file — no Request/Response struct in this file.
	funcFile := `package auth

import "golang.org/x/crypto/bcrypt"

// @func hashPassword
// @description hash

func HashPassword(req HashPasswordRequest) (HashPasswordResponse, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	return HashPasswordResponse{HashedPassword: string(hash)}, nil
}
`
	// Request struct in a separate file.
	reqFile := `package auth

type HashPasswordRequest struct {
	Password string
}
`
	// Response struct in a separate file.
	respFile := `package auth

type HashPasswordResponse struct {
	HashedPassword string
}
`
	os.WriteFile(filepath.Join(authDir, "hash_password.go"), []byte(funcFile), 0644)
	os.WriteFile(filepath.Join(authDir, "hash_password_request.go"), []byte(reqFile), 0644)
	os.WriteFile(filepath.Join(authDir, "hash_password_response.go"), []byte(respFile), 0644)

	specs, diags := ParseDir(dir)
	if len(diags) > 0 {
		t.Fatalf("ParseDir diagnostics: %v", diags)
	}
	if len(specs) != 1 {
		t.Fatalf("ParseDir count = %d, want 1", len(specs))
	}

	spec := specs[0]
	if spec.Name != "hashPassword" {
		t.Errorf("Name = %q, want %q", spec.Name, "hashPassword")
	}
	if len(spec.RequestFields) != 1 {
		t.Fatalf("RequestFields count = %d, want 1", len(spec.RequestFields))
	}
	if spec.RequestFields[0].Name != "Password" {
		t.Errorf("RequestFields[0].Name = %q, want %q", spec.RequestFields[0].Name, "Password")
	}
	if len(spec.ResponseFields) != 1 {
		t.Fatalf("ResponseFields count = %d, want 1", len(spec.ResponseFields))
	}
	if spec.ResponseFields[0].Name != "HashedPassword" {
		t.Errorf("ResponseFields[0].Name = %q, want %q", spec.ResponseFields[0].Name, "HashedPassword")
	}
}
