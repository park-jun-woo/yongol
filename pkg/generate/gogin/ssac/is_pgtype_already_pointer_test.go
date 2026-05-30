//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what isPgtypeAlreadyPointer 단위 테스트 (dotted 필드의 pgtype 변환이 이미 포인터인지)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestIsPgtypeAlreadyPointer(t *testing.T) {
	g := &methodGen{
		VarTypes: map[string]string{
			"user": "User",
			"list": "[]User",
		},
		DDLTables: []ddl.Table{
			{
				Name: "users",
				Columns: map[string]ddl.Column{
					"email": {Name: "email", RawType: "VARCHAR(255)", NotNull: false}, // → *string
					"name":  {Name: "name", RawType: "TEXT", NotNull: true},           // → string
					// BUG-101 acronym guards: leading/inner all-caps fields.
					"id":        {Name: "id", RawType: "UUID", NotNull: true},         // ID → string-equiv value type (not null)
					"id_ref":    {Name: "id_ref", RawType: "UUID", NotNull: false},    // IDRef → *openapi_types.UUID (nullable)
					"parent_id": {Name: "parent_id", RawType: "UUID", NotNull: false}, // ParentID → *openapi_types.UUID (nullable)
				},
			},
		},
	}
	cases := []struct {
		name string
		in   string
		want bool
	}{
		{"no dot", "user", false},
		{"unknown var", "foo.Email", false},
		{"nullable column pointer", "user.Email", true},
		{"not null column non-pointer", "user.Name", false},
		{"slice type stripped, pointer col", "list.Email", true},
		{"missing column", "user.Missing", false},
		// BUG-101: leading all-caps acronym field must resolve to its column.
		// NOT NULL UUID → value type (not pointer) → false.
		{"not null UUID acronym ID", "user.ID", false},
		// nullable UUID with leading acronym → FromPgUUIDPtr returns *T → true.
		{"nullable UUID leading acronym IDRef", "user.IDRef", true},
		// nullable UUID with inner/trailing acronym → *T → true.
		{"nullable UUID inner acronym ParentID", "user.ParentID", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := g.isPgtypeAlreadyPointer(tc.in); got != tc.want {
				t.Errorf("isPgtypeAlreadyPointer(%q) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}
