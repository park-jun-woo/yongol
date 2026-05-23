//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what xdn06ClaimDDLType — 조기 반환 + typed claim ↔ DDL 컬럼 타입 정합/불정합 검증

package manifest_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdn06ClaimDDLType(t *testing.T) {
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
				DDLTables: []ddl.Table{},
			},
			wantCount: 0,
		},
		{
			name: "untyped claims are skipped",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{
							Type:      "jwt",
							UserTable: "users",
							Claims: map[string]pmanifest.ClaimDef{
								"UserID": {Key: "id", GoType: "", Typed: false},
							},
						},
					},
				},
				DDLTables: []ddl.Table{{
					Name: "users",
					Columns: map[string]ddl.Column{
						"id": {RawType: "BIGINT"},
					},
				}},
			},
			wantCount: 0,
		},
		{
			name: "missing column in DDL is skipped (deferred to XDN-03)",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{
							Type:      "jwt",
							UserTable: "users",
							Claims: map[string]pmanifest.ClaimDef{
								"OrgID": {Key: "org_id", GoType: "int64", Typed: true},
							},
						},
					},
				},
				DDLTables: []ddl.Table{{
					Name:    "users",
					Columns: map[string]ddl.Column{},
				}},
			},
			wantCount: 0,
		},
		{
			name: "compatible typed claim returns nil",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{
							Type:      "jwt",
							UserTable: "users",
							Claims: map[string]pmanifest.ClaimDef{
								"UserID": {Key: "id", GoType: "int64", Typed: true},
							},
						},
					},
				},
				DDLTables: []ddl.Table{{
					Name: "users",
					Columns: map[string]ddl.Column{
						"id": {RawType: "BIGINT"},
					},
				}},
			},
			wantCount: 0,
		},
		{
			name: "incompatible typed claim raises XDN-06",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{
							Type:      "jwt",
							UserTable: "users",
							Claims: map[string]pmanifest.ClaimDef{
								"UserID": {Key: "id", GoType: "string", Typed: true},
							},
						},
					},
				},
				DDLTables: []ddl.Table{{
					Name: "users",
					Columns: map[string]ddl.Column{
						"id": {RawType: "BIGINT"},
					},
				}},
			},
			wantCount: 1,
			wantSub:   "XDN-06",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := xdn06ClaimDDLType(tt.fs)
			assertDiagCount(t, diags, tt.wantCount, tt.wantSub)
		})
	}
}
