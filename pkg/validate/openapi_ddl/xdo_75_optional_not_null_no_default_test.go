//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what xdo75OptionalNotNullNoDefault — empty/진단 집계 검증

package openapi_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdo75OptionalNotNullNoDefault(t *testing.T) {
	t.Run("empty constraints returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xdo75OptionalNotNullNoDefault(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("optional NOT NULL field triggers diagnostic", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{
				{Name: "users", Columns: map[string]ddl.Column{
					"name": {NotNull: true, HasDefault: false},
				}},
			},
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"createUser": {"name": {Type: "string", Required: false}},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xdo75OptionalNotNullNoDefault(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
	})
}
