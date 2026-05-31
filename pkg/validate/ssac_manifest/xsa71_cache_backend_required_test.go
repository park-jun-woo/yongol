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

func TestXsa71CacheBackendRequired(t *testing.T) {
	if d := xsa71CacheBackendRequired(nil); d != nil {
		t.Errorf("nil fs -> nil")
	}
	// Not using cache -> nil.
	if d := xsa71CacheBackendRequired(fsWithFuncs()); d != nil {
		t.Errorf("no cache use -> nil, got %v", d)
	}
	usesCacheFn := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("cache.Get")}})
	// Backend declared -> nil.
	usesCacheFn.Manifest = &manifest.ProjectConfig{Cache: &manifest.BuiltinBackend{Backend: "redis"}}
	if d := xsa71CacheBackendRequired(usesCacheFn); d != nil {
		t.Errorf("declared backend -> nil, got %v", d)
	}
	// Backend missing -> ERROR.
	missing := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("cache.Get")}})
	d := xsa71CacheBackendRequired(missing)
	if len(d) != 1 || d[0].Level != diagnostic.LevelError || !strings.Contains(d[0].Message, "[XSA-71]") {
		t.Fatalf("expected XSA-71 error, got %v", d)
	}
}
