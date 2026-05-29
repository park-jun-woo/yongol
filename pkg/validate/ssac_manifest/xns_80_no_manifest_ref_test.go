//ff:func feature=validate type=test control=sequence topic=ssac-manifest
//ff:what XNS-80 참조 없음 — manifest.* 미사용 시 진단 없음

package ssac_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXNS80_NoManifestRef_NoDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{
			{
				Name:     "Login",
				FileName: "auth.go",
				Sequences: []ssac.Sequence{
					{
						Type: "response",
						Line: 10,
						Fields: map[string]string{
							"expires_in": "86400",
						},
					},
				},
			},
		},
	}
	diags := xns80ManifestRef(fs)
	if len(diags) != 0 {
		t.Errorf("expected 0 diagnostics, got %d", len(diags))
	}
}
