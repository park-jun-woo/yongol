//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what Run — nil 안전 + auth 비활성 시 빈 결과 + auth 활성 시 진단 집계 검증

package manifest_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name       string
		fs         *yongol.Fullstack
		wantEmpty  bool
		minDiags   int // minimum number of diagnostics expected (0 means check wantEmpty)
	}{
		{
			name:      "nil fullstack returns empty",
			fs:        nil,
			wantEmpty: true,
		},
		{
			name:      "nil manifest returns empty",
			fs:        &yongol.Fullstack{},
			wantEmpty: true,
		},
		{
			name: "auth absent returns empty",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{},
			},
			wantEmpty: true,
		},
		{
			name: "auth type none returns empty",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{Type: "none"},
					},
				},
			},
			wantEmpty: true,
		},
		{
			name: "auth active but no user_table triggers XDN-01",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{Type: "jwt"},
					},
				},
			},
			wantEmpty: false,
			minDiags:  1,
		},
		{
			name: "auth with user_table but missing DDL triggers XDN-02",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{
							Type:      "jwt",
							UserTable: "users",
						},
					},
				},
			},
			wantEmpty: false,
			minDiags:  1,
		},
		{
			name: "auth with user_table and DDL present, no claims returns clean",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{
							Type:      "jwt",
							UserTable: "users",
						},
					},
				},
				DDLTables: []ddl.Table{
					{Name: "users", Columns: map[string]ddl.Column{}},
				},
			},
			wantEmpty: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := Run(tt.fs)
			if tt.wantEmpty && len(diags) != 0 {
				t.Errorf("expected empty, got %d diagnostics: %+v", len(diags), diags)
			}
			if !tt.wantEmpty && len(diags) < tt.minDiags {
				t.Errorf("expected at least %d diagnostics, got %d", tt.minDiags, len(diags))
			}
		})
	}
}
