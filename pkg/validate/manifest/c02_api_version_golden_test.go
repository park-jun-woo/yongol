//ff:func feature=validate type=test control=sequence topic=manifest-structural
//ff:what c02APIVersion — manifest apiVersion가 "yongol/v1"인지 검증
package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC02APIVersion_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pm.ProjectConfig{APIVersion: "yongol/v1"},
	}
	if got := c02APIVersion(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
