//ff:func feature=validate type=test control=sequence topic=ssac-manifest
//ff:what XNS-80 nil Fullstack — nil 입력 시 진단 없음

package ssac_manifest

import "testing"

func TestXNS80_NilFullstack_NoDiag(t *testing.T) {
	diags := xns80ManifestRef(nil)
	if len(diags) != 0 {
		t.Errorf("expected 0 diagnostics, got %d", len(diags))
	}
}
