//ff:func feature=validate type=test control=sequence
//ff:what TestC09BackendLangFramework_ZeroCov — C-9 백엔드 lang+framework 조합 검증 분기 직접 호출

package manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC09BackendLangFramework_ZeroCov(t *testing.T) {
	// nil fs → nil.
	if d := c09BackendLangFramework(nil); d != nil {
		t.Errorf("nil fs should yield nil")
	}
	// both empty → nil (defaults).
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}
	if d := c09BackendLangFramework(fs); d != nil {
		t.Errorf("empty lang/fw should yield nil, got %v", d)
	}
	// supported combo → nil.
	fs.Manifest.Backend.Lang = "go"
	fs.Manifest.Backend.Framework = "gin"
	if d := c09BackendLangFramework(fs); d != nil {
		t.Errorf("go+gin should be supported, got %v", d)
	}
	// unsupported combo → diagnostic.
	fs.Manifest.Backend.Lang = "ruby"
	fs.Manifest.Backend.Framework = "rails"
	d := c09BackendLangFramework(fs)
	if len(d) != 1 {
		t.Fatalf("unsupported combo should yield 1 diag, got %d", len(d))
	}
}
