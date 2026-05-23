//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what xdo75CheckOpFields — empty/required skip/optional NOT NULL 진단 검증

package openapi_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdo75CheckOpFields(t *testing.T) {
	t.Run("empty fields returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xdo75CheckOpFields(fs, "createUser", map[string]oapiparser.FieldConstraint{})
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("required field skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{
				{Name: "users", Columns: map[string]ddl.Column{
					"email": {NotNull: true},
				}},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		fields := map[string]oapiparser.FieldConstraint{
			"email": {Type: "string", Required: true},
		}
		diags := xdo75CheckOpFields(fs, "createUser", fields)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("optional NOT NULL without DEFAULT raises diagnostic", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{
				{Name: "users", Columns: map[string]ddl.Column{
					"name": {NotNull: true, HasDefault: false},
				}},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		fields := map[string]oapiparser.FieldConstraint{
			"name": {Type: "string", Required: false},
		}
		diags := xdo75CheckOpFields(fs, "createUser", fields)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
	})
}
