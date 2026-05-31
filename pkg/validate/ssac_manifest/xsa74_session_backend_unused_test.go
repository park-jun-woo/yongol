//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what XSA-71/72/74-77 + Run test — backend required(ERROR)/unused(WARNING) 규칙 검증
package ssac_manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXsa74SessionBackendUnused(t *testing.T) {
	if d := xsa74SessionBackendUnused(nil); d != nil {
		t.Errorf("nil fs -> nil")
	}
	// No manifest backend declared -> nil.
	if d := xsa74SessionBackendUnused(fsWithFuncs()); d != nil {
		t.Errorf("no backend declared -> nil, got %v", d)
	}
	// Declared and used -> nil.
	used := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("session.Get")}})
	used.Manifest = &manifest.ProjectConfig{Session: &manifest.BuiltinBackend{Backend: "postgres"}}
	if d := xsa74SessionBackendUnused(used); d != nil {
		t.Errorf("declared+used -> nil, got %v", d)
	}
	// Declared but unused -> WARNING.
	unused := fsWithFuncs()
	unused.Manifest = &manifest.ProjectConfig{Session: &manifest.BuiltinBackend{Backend: "postgres"}}
	d := xsa74SessionBackendUnused(unused)
	if len(d) != 1 || d[0].Level != diagnostic.LevelWarning || !strings.Contains(d[0].Message, "[XSA-74]") {
		t.Fatalf("expected XSA-74 warning, got %v", d)
	}
}
