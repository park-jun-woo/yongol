//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what XSA-71/72/74-77 + Run test — backend required(ERROR)/unused(WARNING) 규칙 검증
package ssac_manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestXsa77QueueBackendUnused(t *testing.T) {
	if d := xsa77QueueBackendUnused(nil); d != nil {
		t.Errorf("nil fs -> nil")
	}
	unused := fsWithFuncs()
	unused.Manifest = &manifest.ProjectConfig{Queue: &manifest.QueueBackend{Backend: "memory"}}
	d := xsa77QueueBackendUnused(unused)
	if len(d) != 1 || d[0].Level != diagnostic.LevelWarning || !strings.Contains(d[0].Message, "[XSA-77]") {
		t.Fatalf("expected XSA-77 warning, got %v", d)
	}
	used := fsWithFuncs(publishFunc())
	used.Manifest = &manifest.ProjectConfig{Queue: &manifest.QueueBackend{Backend: "memory"}}
	if d := xsa77QueueBackendUnused(used); d != nil {
		t.Errorf("declared+used -> nil, got %v", d)
	}
}
