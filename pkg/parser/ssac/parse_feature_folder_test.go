//ff:func feature=ssac-parse type=parser control=sequence
//ff:what feature 서브 폴더 파싱 검증 — Feature 필드가 폴더명으로 설정됨

package ssac

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFeatureFolder(t *testing.T) {
	dir := t.TempDir()
	authDir := filepath.Join(dir, "auth")
	os.MkdirAll(authDir, 0755)

	src := `package service

// @get User user = User.FindByEmail({Email: request.Email})
// @response {
//   user: user
// }
func Login(c *gin.Context) {}
`
	os.WriteFile(filepath.Join(authDir, "login.ssac"), []byte(src), 0644)

	funcs, diags := ParseDir(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diagnostics: %v", diags[0].Message)
	}
	if len(funcs) != 1 {
		t.Fatalf("expected 1 func, got %d", len(funcs))
	}
	assertEqual(t, "Feature", funcs[0].Feature, "auth")
}
