//ff:func feature=validate type=test control=sequence topic=ssac-manifest
//ff:what XNS-80 미지원 경로 — 에러 진단 검증

package ssac_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXNS80_UnknownPath_Error(t *testing.T) {
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
							"expires_in": "manifest.auth.unknownField",
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
	if diags[0].Level != diagnostic.LevelError {
		t.Errorf("expected error level, got %s", diags[0].Level)
	}
}
