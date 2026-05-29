//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what xdo69CheckValuesEnum — 일치/불일치/한쪽 누락 검증

package openapi_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdo69CheckValuesEnum_Unit(t *testing.T) {
	t.Run("empty constraints returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xdo69CheckValuesEnum(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("matching enum values passes", func(t *testing.T) {
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
		diags := xdo69CheckValuesEnum(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("mismatched enum values raises diagnostic", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{
				{Name: "users", Columns: map[string]ddl.Column{
					"role": {CheckEnum: []string{"admin", "user"}},
				}},
			},
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"createUser": {"role": {Type: "string", Enum: []string{"admin", "moderator"}}},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xdo69CheckValuesEnum(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "XDO-69") {
			t.Errorf("Message missing XDO-69: %s", diags[0].Message)
		}
	})

	t.Run("no CHECK skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{
				{Name: "users", Columns: map[string]ddl.Column{"role": {}}},
			},
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"createUser": {"role": {Type: "string", Enum: []string{"admin"}}},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xdo69CheckValuesEnum(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}

func TestXdo69CheckValuesEnum(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
