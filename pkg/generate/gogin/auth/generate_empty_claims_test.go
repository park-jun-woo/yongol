//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerateSkipsWhenEmptyClaims — auth 존재하나 claims 비면 생성 안 함
package auth

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerateSkipsWhenEmptyClaims(t *testing.T) {
	dir := t.TempDir()
	fs := &yongol.Fullstack{}
	p := prepared.State{Auth: prepared.Auth{
		Present: true,
		Raw:     &manifest.Auth{Claims: map[string]manifest.ClaimDef{}},
	}}
	if err := Generate(fs, p, dir); err != nil {
		t.Fatalf("Generate error: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "backend")); !os.IsNotExist(err) {
		t.Errorf("expected no output for empty claims, stat err: %v", err)
	}
}
