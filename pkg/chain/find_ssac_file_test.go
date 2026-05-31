//ff:func feature=chain type=test control=sequence
//ff:what findSSaCFile 가 feature 폴더 구조 / flat 구조 / 미존재 fallback 을 올바르게 선택하는지 검증
package chain

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestFindSSaCFile(t *testing.T) {
	specsDir := t.TempDir()

	// Feature-folder structure present.
	featDir := filepath.Join(specsDir, "service", "auth")
	if err := os.MkdirAll(featDir, 0o755); err != nil {
		t.Fatalf("mkdir feat: %v", err)
	}
	if err := os.WriteFile(filepath.Join(featDir, "login.ssac"), []byte("x"), 0o644); err != nil {
		t.Fatalf("write feat: %v", err)
	}
	sf := &ssac.ServiceFunc{Name: "Login", FileName: "login.ssac", Feature: "auth"}
	if got := findSSaCFile(sf, specsDir); got != filepath.Join("service", "auth", "login.ssac") {
		t.Errorf("feature folder: got %q", got)
	}

	// Flat structure: feature set but file only exists flat.
	flatFile := filepath.Join(specsDir, "service", "signup.ssac")
	if err := os.WriteFile(flatFile, []byte("x"), 0o644); err != nil {
		t.Fatalf("write flat: %v", err)
	}
	sfFlat := &ssac.ServiceFunc{Name: "Signup", FileName: "signup.ssac", Feature: "missing"}
	if got := findSSaCFile(sfFlat, specsDir); got != filepath.Join("service", "signup.ssac") {
		t.Errorf("flat structure: got %q", got)
	}

	// Neither exists → fallback path.
	sfNone := &ssac.ServiceFunc{Name: "Nope", FileName: "nope.ssac"}
	if got := findSSaCFile(sfNone, specsDir); got != "service/nope.ssac" {
		t.Errorf("fallback: got %q", got)
	}
}
