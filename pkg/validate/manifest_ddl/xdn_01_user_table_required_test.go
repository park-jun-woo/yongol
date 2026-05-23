//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what xdn01UserTableRequired — auth 비활성 조기 반환 + user_table 유무 검증

package manifest_ddl

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdn01UserTableRequired(t *testing.T) {
	tests := []struct {
		name      string
		fs        *yongol.Fullstack
		wantDiags bool
		wantSub   string
	}{
		{
			name:      "nil fullstack returns nil",
			fs:        nil,
			wantDiags: false,
		},
		{
			name: "auth absent returns nil",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{},
			},
			wantDiags: false,
		},
		{
			name: "auth type none returns nil",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{Type: "none"},
					},
				},
			},
			wantDiags: false,
		},
		{
			name: "auth active with user_table set returns nil",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{Type: "jwt", UserTable: "users"},
					},
				},
			},
			wantDiags: false,
		},
		{
			name: "auth active without user_table raises XDN-01",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{Type: "jwt"},
					},
				},
			},
			wantDiags: true,
			wantSub:   "XDN-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := xdn01UserTableRequired(tt.fs)
			assertDiags(t, diags, tt.wantDiags, tt.wantSub)
		})
	}
}
