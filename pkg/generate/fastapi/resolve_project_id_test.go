//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestFastapiHelpers — fastapi plan/package/route 헬퍼 검증 (Op 종류·외부 패키지 수집·라우트 해석)
package fastapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveProjectID(t *testing.T) {
	// No manifest -> "app".
	if got := resolveProjectID(&yongol.Fullstack{}); got != "app" {
		t.Errorf("expected app fallback, got %q", got)
	}
	// Manifest with name.
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{
		Metadata: manifest.Metadata{Name: "zenflow"},
	}}
	if got := resolveProjectID(fs); got != "zenflow" {
		t.Errorf("expected zenflow, got %q", got)
	}
	// Manifest with empty name -> fallback.
	fs2 := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	if got := resolveProjectID(fs2); got != "app" {
		t.Errorf("expected app for empty name, got %q", got)
	}
}
