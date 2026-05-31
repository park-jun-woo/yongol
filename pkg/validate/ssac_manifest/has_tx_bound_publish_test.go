//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what TestSSaCManifestHelpers — unit tests for the pure ssac_manifest helper functions
package ssac_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestHasTxBoundPublish(t *testing.T) {
	// mutation + publish → true.
	fn := ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "post"}, {Type: "publish"}}}
	if !hasTxBoundPublish(fn) {
		t.Error("post + publish should be tx-bound")
	}
	// publish without mutation → false.
	if hasTxBoundPublish(ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "get"}, {Type: "publish"}}}) {
		t.Error("publish w/o mutation should be false")
	}
	// mutation without publish → false.
	if hasTxBoundPublish(ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "put"}}}) {
		t.Error("mutation w/o publish should be false")
	}
}
