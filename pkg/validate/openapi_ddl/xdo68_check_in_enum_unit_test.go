//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what xdo68CheckInEnum — empty/CHECK 없음/enum 유무 검증
package openapi_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdo68CheckInEnum_Unit(t *testing.T) {
	t.Run("empty constraints returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xdo68CheckInEnum(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no CHECK constraint skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{
				{Name: "users", Columns: map[string]ddl.Column{"email": {RawType: "TEXT"}}},
			},
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"createUser": {"email": {Type: "string"}},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xdo68CheckInEnum(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("CHECK with enum present passes", func(t *testing.T) {
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
		diags := xdo68CheckInEnum(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("CHECK without enum raises diagnostic", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{
				{Name: "users", Columns: map[string]ddl.Column{
					"role": {CheckEnum: []string{"admin", "user"}},
				}},
			},
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"createUser": {"role": {Type: "string"}},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xdo68CheckInEnum(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "XDO-68") {
			t.Errorf("Message missing XDO-68: %s", diags[0].Message)
		}
	})
}
