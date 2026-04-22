//ff:func feature=funcspec type=parser control=sequence
//ff:what FuncSpec.Line 이 @func 어노테이션 줄 번호로 채워지는지 검증

package funcspec

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseLine_FuncSpecAnnotation(t *testing.T) {
	dir := t.TempDir()
	pkgDir := filepath.Join(dir, "auth")
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	// Lines:
	// 1: package auth
	// 2: blank
	// 3: // @func hashPassword
	// 4: // @description ...
	// 5: blank
	// 6: type HashPasswordRequest struct {
	// ...
	src := `package auth

// @func hashPassword
// @description 평문 비밀번호를 bcrypt 해시로 변환한다

type HashPasswordRequest struct {
	Password string
}

type HashPasswordResponse struct {
	HashedPassword string
}

func HashPassword(req HashPasswordRequest) (HashPasswordResponse, error) {
	return HashPasswordResponse{}, nil
}
`
	path := filepath.Join(pkgDir, "hash_password.go")
	if err := os.WriteFile(path, []byte(src), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}

	spec, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("ParseFile diagnostics: %v", diags)
	}
	if spec == nil {
		t.Fatal("expected non-nil spec")
	}

	if spec.Line != 3 {
		t.Errorf("FuncSpec.Line = %d, want 3 (@func annotation line)", spec.Line)
	}
}
