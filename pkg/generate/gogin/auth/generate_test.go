//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerateSkipsWhenNoAuth — auth 미존재 시 아무것도 생성하지 않음
package auth

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerateSkipsWhenNoAuth(t *testing.T) {
	dir := t.TempDir()
	fs := &yongol.Fullstack{}
	p := prepared.State{Auth: prepared.Auth{Present: false}}
	if err := Generate(fs, p, dir); err != nil {
		t.Fatalf("Generate error: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "backend")); !os.IsNotExist(err) {
		t.Errorf("expected no output for absent auth, stat err: %v", err)
	}
}
