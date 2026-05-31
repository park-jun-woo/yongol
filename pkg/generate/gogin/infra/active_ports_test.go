//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestActivePorts — when 이 truthy 한 port 만 남기는지 검증
package infra

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

func TestActivePorts(t *testing.T) {
	ports := []ssacmeta.Port{
		{Name: "always-on", When: ""},                        // truthy -> kept
		{Name: "explicit-always", When: "always"},            // truthy -> kept
		{Name: "off", When: "manifest.cache.backend=='off'"}, // false under empty mctx -> dropped
	}
	mctx := map[string]any{}
	got := activePorts(ports, mctx)

	if len(got) != 2 {
		t.Fatalf("expected 2 active ports, got %d: %v", len(got), got)
	}
	names := map[string]bool{}
	for _, p := range got {
		names[p.Name] = true
	}
	if !names["always-on"] || !names["explicit-always"] {
		t.Errorf("expected truthy ports kept, got: %v", names)
	}
	if names["off"] {
		t.Errorf("inactive port should be dropped")
	}
}
