//ff:func feature=validate type=test control=sequence topic=manifest-structural
//ff:what C-2 테스트 — APIVersion golden

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestC02APIVersion_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{APIVersion: "yongol/v1"},
	}
	if got := c02APIVersion(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
