//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what xna90RefreshRequiresSQLC — nil/absent 조기 반환 + auth 존재 시 위임 검증

package manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXna90RefreshRequiresSQLC(t *testing.T) {
	tests := []struct {
		name      string
		fs        *yongol.Fullstack
		wantDiags bool
		wantSub   string // substring to look for in message when wantDiags=true
	}{
		{
			name:      "nil fullstack returns nil",
			fs:        nil,
			wantDiags: false,
		},
		{
			name:      "nil manifest returns nil",
			fs:        &yongol.Fullstack{},
			wantDiags: false,
		},
		{
			name: "nil auth returns nil",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{},
			},
			wantDiags: false,
		},
		{
			name: "auth present but no DDL/sqlc raises diagnostic",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{SecretEnv: "JWT_SECRET"},
					},
				},
			},
			wantDiags: true,
			wantSub:   "XNA-90",
		},
		{
			name: "diagnostic mentions refresh_tokens",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{SecretEnv: "JWT_SECRET"},
					},
				},
			},
			wantDiags: true,
			wantSub:   "refresh_tokens",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := xna90RefreshRequiresSQLC(tt.fs)
			assertDiags(t, diags, tt.wantDiags, tt.wantSub)
		})
	}
}
