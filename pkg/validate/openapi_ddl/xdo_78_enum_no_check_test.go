//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what xdo78EnumNoCheck — enum 없음/CHECK 있음/CHECK 없음 검증

package openapi_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdo78EnumNoCheck(t *testing.T) {
	t.Run("empty constraints returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xdo78EnumNoCheck(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no enum in OpenAPI skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{
				{Name: "users", Columns: map[string]ddl.Column{"role": {}}},
			},
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"createUser": {"role": {Type: "string"}},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xdo78EnumNoCheck(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("enum with CHECK present passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{
				{Name: "users", Columns: map[string]ddl.Column{
					"role": {CheckEnum: []string{"admin", "user"}},
				}},
			},
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"createUser": {"role": {Type: "string", Enum: []string{"admin", "user"}}},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xdo78EnumNoCheck(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("column not in DDL skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{},
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"createUser": {"role": {Type: "string", Enum: []string{"admin"}}},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xdo78EnumNoCheck(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("enum without CHECK raises diagnostic", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{
				{Name: "users", Columns: map[string]ddl.Column{
					"role": {},
				}},
			},
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"createUser": {"role": {Type: "string", Enum: []string{"admin", "user"}}},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xdo78EnumNoCheck(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "XDO-78") {
			t.Errorf("Message missing XDO-78: %s", diags[0].Message)
		}
	})
}
