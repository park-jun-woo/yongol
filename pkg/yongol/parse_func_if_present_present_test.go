//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseFuncIfPresent — Func 미탐지(return) + 탐지 시 ProjectFuncSpecs 설정
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFuncIfPresent_Present(t *testing.T) {
	dir := t.TempDir()
	authDir := filepath.Join(dir, "auth")
	if err := os.MkdirAll(authDir, 0o755); err != nil {
		t.Fatal(err)
	}
	funcFile := `package auth

// @func hashPassword
// @description hash

func HashPassword(req HashPasswordRequest) (HashPasswordResponse, error) {
	return HashPasswordResponse{}, nil
}

type HashPasswordRequest struct {
	Password string
}

type HashPasswordResponse struct {
	HashedPassword string
}
`
	if err := os.WriteFile(filepath.Join(authDir, "hash_password.go"), []byte(funcFile), 0o644); err != nil {
		t.Fatal(err)
	}

	fs := &Fullstack{}
	has := map[SSOTKind]DetectedSSOT{
		KindFunc: {Kind: KindFunc, Path: dir, Presence: SSOTPopulated},
	}
	parseFuncIfPresent(fs, has)
	if len(fs.ProjectFuncSpecs) == 0 {
		t.Fatalf("expected ProjectFuncSpecs populated, diags=%+v", fs.ParseDiagnostics)
	}
}
