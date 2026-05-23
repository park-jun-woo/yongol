//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what xdn02UserTableExists — auth 비활성/empty user_table/DDL 유무 검증

package manifest_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdn02UserTableExists(t *testing.T) {
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
			name: "auth inactive returns nil",
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
			name: "empty user_table returns nil (XDN-01 covers this)",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{Type: "jwt", UserTable: ""},
					},
				},
			},
			wantDiags: false,
		},
		{
			name: "user_table matches DDL table returns nil",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{Type: "jwt", UserTable: "users"},
					},
				},
				DDLTables: []ddl.Table{{Name: "users"}},
			},
			wantDiags: false,
		},
		{
			name: "user_table does not match any DDL table raises XDN-02",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{Type: "jwt", UserTable: "users"},
					},
				},
				DDLTables: []ddl.Table{{Name: "orders"}},
			},
			wantDiags: true,
			wantSub:   "XDN-02",
		},
		{
			name: "diagnostic mentions the missing table name",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{Type: "jwt", UserTable: "accounts"},
					},
				},
			},
			wantDiags: true,
			wantSub:   "accounts",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := xdn02UserTableExists(tt.fs)
			assertDiags(t, diags, tt.wantDiags, tt.wantSub)
		})
	}
}
