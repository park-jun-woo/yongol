//ff:func feature=validate-contract type=test control=sequence
//ff:what TestResolveImportName — import alias 또는 경로 basename 으로 패키지 식별자 결정 검증

package contract

import (
	"go/ast"
	"testing"
)

func TestResolveImportName(t *testing.T) {
	tests := []struct {
		name  string
		alias string
		path  string
		want  string
	}{
		{"no alias basename", "", "\"database/sql\"", "sql"},
		{"single segment", "", "\"fmt\"", "fmt"},
		{"explicit alias", "sqlx", "\"github.com/jmoiron/sqlx\"", "sqlx"},
		{"blank alias falls back", "_", "\"database/sql\"", "sql"},
		{"dot alias falls back", ".", "\"fmt\"", "fmt"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ident *ast.Ident
			if tt.alias != "" {
				ident = &ast.Ident{Name: tt.alias}
			}
			if got := resolveImportName(ident, tt.path); got != tt.want {
				t.Fatalf("resolveImportName(%q, %q) = %q, want %q", tt.alias, tt.path, got, tt.want)
			}
		})
	}
}
