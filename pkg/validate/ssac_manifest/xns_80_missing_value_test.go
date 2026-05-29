//ff:func feature=validate type=test control=sequence topic=ssac-manifest
//ff:what XNS-80 값 누락 — 경로 유효하지만 값 비어 있으면 에러

package ssac_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXNS80_KnownPathMissingValue_Error(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				Auth: &manifest.Auth{},
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
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diags))
	}
}
