//ff:func feature=validate type=test control=sequence topic=manifest-structural
//ff:what c04MetadataName — manifest metadata.name 비어있는지 검증
package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC04MetadataName_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pm.ProjectConfig{
			Metadata: pm.Metadata{Name: "zenflow"},
		},
	}
	if got := c04MetadataName(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
