//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what XNS-56 test — @publish/@subscribe requires manifest queue.backend (Ground.Config)

package ssac_manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestXns56QueueRequired(t *testing.T) {
	// nil fs -> nil.
	if d := xns56QueueRequired(nil); d != nil {
		t.Errorf("nil fs should yield nil, got %v", d)
	}

	// No queue use -> nil.
	noQueue := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("cache.Get")}})
	if d := xns56QueueRequired(noQueue); d != nil {
		t.Errorf("no queue use should yield nil, got %v", d)
	}

	// Uses publish + queue.backend configured in Ground -> nil.
	pubFn := ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "publish"}}}
	configured := fsWithFuncs(pubFn)
	configured.SetGround(&rule.Ground{Config: map[string]bool{"queue.backend": true}})
	if d := xns56QueueRequired(configured); d != nil {
		t.Errorf("configured queue should yield nil, got %v", d)
	}

	// Uses publish but no queue.backend -> diagnostic.
	missing := fsWithFuncs(pubFn)
	d := xns56QueueRequired(missing)
	if len(d) != 1 || !strings.Contains(d[0].Message, "[XNS-56]") {
		t.Fatalf("expected XNS-56 diag, got %v", d)
	}
}
