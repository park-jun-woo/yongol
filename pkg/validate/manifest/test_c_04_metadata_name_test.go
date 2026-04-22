//ff:func feature=validate type=test control=sequence topic=manifest-structural
//ff:what C-4 테스트 — metadata.name golden

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestC04MetadataName_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Metadata: pmanifest.Metadata{Name: "zenflow"},
		},
	}
	if got := c04MetadataName(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
