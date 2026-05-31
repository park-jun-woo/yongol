//ff:func feature=validate type=test control=iteration dimension=1 topic=config-check
//ff:what XSA-71/72/74-77 + Run test — backend required(ERROR)/unused(WARNING) 규칙 검증
package ssac_manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun(t *testing.T) {
	// Empty Fullstack -> no diagnostics (all rules short-circuit).
	if d := Run(&yongol.Fullstack{}); len(d) != 0 {
		t.Errorf("empty fs Run should yield no diags, got %v", d)
	}
	// A single XSA-75 warning aggregated through Run.
	fs := fsWithFuncs()
	fs.Manifest = &manifest.ProjectConfig{Cache: &manifest.BuiltinBackend{Backend: "redis"}}
	d := Run(fs)
	found := false
	for _, x := range d {
		if strings.Contains(x.Message, "[XSA-75]") {
			found = true
		}
	}
	if !found {
		t.Errorf("expected XSA-75 in aggregated Run output, got %v", d)
	}
}
