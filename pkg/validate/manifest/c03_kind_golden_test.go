//ff:func feature=validate type=test control=sequence topic=manifest-structural
//ff:what c03Kind — manifest kind가 "Project"인지 검증
package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC03Kind_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pm.ProjectConfig{Kind: "Project"},
	}
	if got := c03Kind(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
