//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what TestSSaCManifestHelpers — unit tests for the pure ssac_manifest helper functions
package ssac_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestUsesSession(t *testing.T) {
	yes := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("session.Put")}})
	if !usesSession(yes) {
		t.Error("expected session use detected")
	}
	no := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("cache.Get")}})
	if usesSession(no) {
		t.Error("cache.Get should not count as session use")
	}
}
