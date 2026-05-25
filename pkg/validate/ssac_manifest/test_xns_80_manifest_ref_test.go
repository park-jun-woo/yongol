//ff:func feature=validate type=test control=sequence topic=ssac-manifest
//ff:what XNS-80 유효 참조 — 진단 없음 검증

package ssac_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXNS80_ValidRef_NoDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				Auth: &manifest.Auth{AccessTokenTTL: "15m"},
			},
		},
		ServiceFuncs: []ssac.ServiceFunc{
			{
				Name:     "Login",
				FileName: "auth.go",
				Sequences: []ssac.Sequence{
					{
						Type: "response",
						Line: 10,
						Fields: map[string]string{
							"expires_in": "manifest.auth.accessTokenTTL",
						},
					},
				},
			},
		},
	}
	diags := xns80ManifestRef(fs)
	if len(diags) != 0 {
		t.Errorf("expected 0 diagnostics, got %d: %v", len(diags), diags)
	}
}
