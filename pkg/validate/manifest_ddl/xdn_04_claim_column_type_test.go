//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what xdn04ClaimColumnType — 조기 반환 + 타입 일치/불일치 진단 개수 검증

package manifest_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdn04ClaimColumnType(t *testing.T) {
	tests := []struct {
		name      string
		fs        *yongol.Fullstack
		wantCount int
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
			name: "all claims type-compatible returns nil",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{
							Type:      "jwt",
							UserTable: "users",
							Claims: map[string]pmanifest.ClaimDef{
								"UserID": {Key: "id", GoType: "int64"},
								"Email":  {Key: "email", GoType: "string"},
							},
						},
					},
				},
				DDLTables: []ddl.Table{{
					Name: "users",
					Columns: map[string]ddl.Column{
						"id":    {RawType: "BIGINT"},
						"email": {RawType: "TEXT"},
					},
				}},
			},
			wantCount: 0,
		},
		{
			name: "one type mismatch raises one diagnostic",
			fs: &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Auth: &pmanifest.Auth{
							Type:      "jwt",
							UserTable: "users",
							Claims: map[string]pmanifest.ClaimDef{
								"UserID": {Key: "id", GoType: "string"},
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := xdn04ClaimColumnType(tt.fs)
			if len(diags) != tt.wantCount {
				t.Errorf("expected %d diagnostics, got %d: %+v", tt.wantCount, len(diags), diags)
			}
		})
	}
}
