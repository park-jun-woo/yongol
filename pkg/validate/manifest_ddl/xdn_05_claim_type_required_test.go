//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what xdn05ClaimTypeRequired — 비활성/빈 claims 조기 반환 + typed/untyped/invalid 타입 검증

package manifest_ddl

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdn05ClaimTypeRequired(t *testing.T) {
	tests := []struct {
		name      string
		fs        *yongol.Fullstack
		wantCount int
		wantSub   string
	}{
		{
			name:      "nil fullstack returns nil",
			fs:        nil,
			wantCount: 0,
		},
		{
			name: "auth inactive returns nil",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{},
			},
			wantCount: 0,
		},
		{
			name: "no claims returns nil",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{Type: "jwt"},
					},
				},
			},
			wantCount: 0,
		},
		{
			name: "all claims typed with allowed types returns nil",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{
							Type: "jwt",
							Claims: map[string]pmanifest.ClaimDef{
								"UserID": {Key: "id", GoType: "int64", Typed: true},
								"Email":  {Key: "email", GoType: "string", Typed: true},
							},
						},
					},
				},
			},
			wantCount: 0,
		},
		{
			name: "untyped claim raises diagnostic",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{
							Type: "jwt",
							Claims: map[string]pmanifest.ClaimDef{
								"UserID": {Key: "id", GoType: "", Typed: false},
							},
						},
					},
				},
			},
			wantCount: 1,
			wantSub:   "XDN-05",
		},
		{
			name: "invalid type raises diagnostic",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{
							Type: "jwt",
							Claims: map[string]pmanifest.ClaimDef{
								"Data": {Key: "data", GoType: "float64", Typed: true},
							},
						},
					},
				},
			},
			wantCount: 1,
			wantSub:   "XDN-05",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := xdn05ClaimTypeRequired(tt.fs)
			assertDiagCount(t, diags, tt.wantCount, tt.wantSub)
		})
	}
}
