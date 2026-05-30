//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseSSaCIfPresent — SSaC 미탐지(return) + 탐지 시 ServiceFuncs 설정

package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseSSaCIfPresent_Absent(t *testing.T) {
	fs := &Fullstack{}
	parseSSaCIfPresent(fs, map[SSOTKind]DetectedSSOT{})
	if fs.ServiceFuncs != nil {
		t.Fatalf("expected no ServiceFuncs when absent")
	}
}

func TestParseSSaCIfPresent_Present(t *testing.T) {
	dir := t.TempDir()
	authDir := filepath.Join(dir, "auth")
	if err := os.MkdirAll(authDir, 0o755); err != nil {
		t.Fatal(err)
	}
	src := `package service

// @get User user = User.FindByEmail({Email: request.Email})
// @response {
//   user: user
// }
func Login(c *gin.Context) {}
`
	if err := os.WriteFile(filepath.Join(authDir, "login.ssac"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	fs := &Fullstack{}
	has := map[SSOTKind]DetectedSSOT{
		KindSSaC: {Kind: KindSSaC, Path: dir, Presence: SSOTPopulated},
	}
	parseSSaCIfPresent(fs, has)
	if len(fs.ServiceFuncs) == 0 {
		t.Fatalf("expected ServiceFuncs populated, diags=%+v", fs.ParseDiagnostics)
	}
}
