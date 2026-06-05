//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestCollectComponentApiOpsMissingFileSkipped — .tsx 부재 시 소비 0 (skip)

package stml_openapi

import "testing"

func TestCollectComponentApiOpsMissingFileSkipped(t *testing.T) {
	specsDir := t.TempDir()
	names := map[string]struct{}{"Absent": {}}
	ops := map[string]struct{}{"AnyOp": {}}
	out := make(map[string]struct{})

	collectComponentApiOps(names, specsDir, ops, out)

	if len(out) != 0 {
		t.Errorf("missing file should yield no consumption, got %+v", out)
	}
}
