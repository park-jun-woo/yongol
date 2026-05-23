//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what xdo76RequiredNullable — optional skip/NOT NULL skip/@nullable skip/진단 검증

package openapi_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdo76RequiredNullable(t *testing.T) {
	t.Run("empty constraints returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xdo76RequiredNullable(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("optional field skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{
				{Name: "users", Columns: map[string]ddl.Column{"bio": {}}},
			},
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"createUser": {"bio": {Required: false}},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xdo76RequiredNullable(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("required field with no DDL table skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{},
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"createUser": {"ghost_field": {Required: true}},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xdo76RequiredNullable(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("required NOT NULL column passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{
				{Name: "users", Columns: map[string]ddl.Column{
					"email": {NotNull: true},
				}},
			},
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"createUser": {"email": {Required: true}},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xdo76RequiredNullable(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("required nullable with @nullable annotation exempt", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{
				{Name: "users", Columns: map[string]ddl.Column{
					"avatar": {NullableAnnot: true},
				}},
			},
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"createUser": {"avatar": {Required: true}},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xdo76RequiredNullable(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("required nullable without annotation raises warning", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DDLTables: []ddl.Table{
				{Name: "users", Columns: map[string]ddl.Column{
					"bio": {},
				}},
			},
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"createUser": {"bio": {Required: true}},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xdo76RequiredNullable(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "XDO-76") {
			t.Errorf("Message missing XDO-76: %s", diags[0].Message)
		}
	})
}
