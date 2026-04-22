//ff:func feature=validate type=test control=sequence topic=manifest-structural
//ff:what C-3 테스트 — Kind golden

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestC03Kind_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{Kind: "Project"},
	}
	if got := c03Kind(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
