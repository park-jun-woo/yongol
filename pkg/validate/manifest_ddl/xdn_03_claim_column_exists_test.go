//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what xdn03ClaimColumnExists — 조기 반환 + claim column 존재/누락 검증

package manifest_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdn03ClaimColumnExists(t *testing.T) {
	tests := []struct {
		name      string
		fs        *yongol.Fullstack
		wantCount int
		wantSub   string // substring in first diag message
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
			name: "empty user_table returns nil",
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
			name: "user_table not in DDL returns nil",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{Type: "jwt", UserTable: "users"},
					},
				},
				DDLTables: []ddl.Table{{Name: "orders"}},
			},
			wantCount: 0,
		},
		{
			name: "all claim columns exist returns nil",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{
							Type:      "jwt",
							UserTable: "users",
							Claims: map[string]pmanifest.ClaimDef{
								"UserID": {Key: "id"},
								"Email":  {Key: "email"},
							},
						},
					},
				},
				DDLTables: []ddl.Table{{
					Name: "users",
					Columns: map[string]ddl.Column{
						"id":    {},
						"email": {},
					},
				}},
			},
			wantCount: 0,
		},
		{
			name: "missing claim column raises one diagnostic",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{
							Type:      "jwt",
							UserTable: "users",
							Claims: map[string]pmanifest.ClaimDef{
								"OrgID": {Key: "org_id"},
							},
						},
					},
				},
				DDLTables: []ddl.Table{{
					Name:    "users",
					Columns: map[string]ddl.Column{"id": {}},
				}},
			},
			wantCount: 1,
			wantSub:   "XDN-03",
		},
		{
			name: "two missing columns raise two diagnostics",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{
							Type:      "jwt",
							UserTable: "users",
							Claims: map[string]pmanifest.ClaimDef{
								"OrgID": {Key: "org_id"},
								"Role":  {Key: "role"},
							},
						},
					},
				},
				DDLTables: []ddl.Table{{
					Name:    "users",
					Columns: map[string]ddl.Column{"id": {}},
				}},
			},
			wantCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := xdn03ClaimColumnExists(tt.fs)
			assertDiagCount(t, diags, tt.wantCount, tt.wantSub)
		})
	}
}
