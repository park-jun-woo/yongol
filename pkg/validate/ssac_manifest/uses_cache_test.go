//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what TestSSaCManifestHelpers — unit tests for the pure ssac_manifest helper functions
package ssac_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestUsesCache(t *testing.T) {
	yes := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("cache.Get")}})
	if !usesCache(yes) {
		t.Error("expected cache use detected")
	}
	no := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("session.Get")}})
	if usesCache(no) {
		t.Error("session.Get should not count as cache use")
	}
	if usesCache(fsWithFuncs()) {
		t.Error("empty fs should not use cache")
	}
}
