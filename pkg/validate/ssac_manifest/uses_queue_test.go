//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what TestSSaCManifestHelpers — unit tests for the pure ssac_manifest helper functions
package ssac_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestUsesQueue(t *testing.T) {
	// publish sequence.
	if !usesQueue(fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "publish"}}})) {
		t.Error("publish should count")
	}
	// subscribe func.
	if !usesQueue(fsWithFuncs(ssac.ServiceFunc{Subscribe: &ssac.SubscribeInfo{}})) {
		t.Error("subscribe should count")
	}
	if usesQueue(fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("cache.Get")}})) {
		t.Error("no queue usage expected")
	}
}
