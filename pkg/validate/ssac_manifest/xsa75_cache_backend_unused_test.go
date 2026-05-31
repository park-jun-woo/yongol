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

func TestXsa75CacheBackendUnused(t *testing.T) {
	if d := xsa75CacheBackendUnused(nil); d != nil {
		t.Errorf("nil fs -> nil")
	}
	unused := fsWithFuncs()
	unused.Manifest = &manifest.ProjectConfig{Cache: &manifest.BuiltinBackend{Backend: "redis"}}
	d := xsa75CacheBackendUnused(unused)
	if len(d) != 1 || d[0].Level != diagnostic.LevelWarning || !strings.Contains(d[0].Message, "[XSA-75]") {
		t.Fatalf("expected XSA-75 warning, got %v", d)
	}
	used := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("cache.Set")}})
	used.Manifest = &manifest.ProjectConfig{Cache: &manifest.BuiltinBackend{Backend: "redis"}}
	if d := xsa75CacheBackendUnused(used); d != nil {
		t.Errorf("declared+used -> nil, got %v", d)
	}
	if d := xsa75CacheBackendUnused(fsWithFuncs()); d != nil {
		t.Errorf("no backend -> nil, got %v", d)
	}
}
