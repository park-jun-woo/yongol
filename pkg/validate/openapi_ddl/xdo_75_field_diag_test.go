//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what xdo75FieldDiag — required/nullable/NOT NULL+DEFAULT/NOT NULL 진단 검증

package openapi_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXdo75FieldDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{
			{
				Name: "users",
				Columns: map[string]ddl.Column{
					"name": {NotNull: true, HasDefault: false},
					"role": {NotNull: true, HasDefault: true},
					"bio":  {NotNull: false},
				},
			},
		},
		OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
	}

	t.Run("required field returns false", func(t *testing.T) {
		_, ok := xdo75FieldDiag(fs, "createUser", "name", oapiparser.FieldConstraint{Required: true})
		if ok {
			t.Error("expected false for required field")
		}
	})

	t.Run("nullable column returns false", func(t *testing.T) {
		_, ok := xdo75FieldDiag(fs, "createUser", "bio", oapiparser.FieldConstraint{})
		if ok {
			t.Error("expected false for nullable column")
		}
	})

	t.Run("NOT NULL with DEFAULT returns false", func(t *testing.T) {
		_, ok := xdo75FieldDiag(fs, "createUser", "role", oapiparser.FieldConstraint{})
		if ok {
			t.Error("expected false for NOT NULL with DEFAULT")
		}
	})

	t.Run("NOT NULL without DEFAULT returns diagnostic", func(t *testing.T) {
		diag, ok := xdo75FieldDiag(fs, "createUser", "name", oapiparser.FieldConstraint{})
		if !ok {
			t.Fatal("expected true for NOT NULL without DEFAULT")
		}
		if !strings.Contains(diag.Message, "XDO-75") {
			t.Errorf("Message missing XDO-75: %s", diag.Message)
		}
	})

	t.Run("column not in DDL returns false", func(t *testing.T) {
		_, ok := xdo75FieldDiag(fs, "createUser", "ghost", oapiparser.FieldConstraint{})
		if ok {
			t.Error("expected false for missing column")
		}
	})
}
