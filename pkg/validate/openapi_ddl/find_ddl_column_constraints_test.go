//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what findDDLColumnConstraints — 컬럼 미발견/발견 + VARCHAR/CHECK 추출 검증

package openapi_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFindDDLColumnConstraints(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{
			{
				Name: "users",
				Columns: map[string]ddl.Column{
					"email": {RawType: "VARCHAR(255)", VarcharLen: 255},
					"role":  {RawType: "VARCHAR(20)", VarcharLen: 20, CheckEnum: []string{"admin", "user"}},
					"id":    {RawType: "BIGINT"},
				},
			},
		},
	}

	t.Run("column not found", func(t *testing.T) {
		_, _, found := findDDLColumnConstraints(fs, "nonexistent")
		if found {
			t.Error("expected found=false for nonexistent column")
		}
	})

	t.Run("column without constraints", func(t *testing.T) {
		vlen, enums, found := findDDLColumnConstraints(fs, "id")
		if !found {
			t.Fatal("expected found=true")
		}
		if vlen != 0 {
			t.Errorf("VarcharLen = %d, want 0", vlen)
		}
		if len(enums) != 0 {
			t.Errorf("CheckEnum = %v, want empty", enums)
		}
	})

	t.Run("VARCHAR length extracted", func(t *testing.T) {
		vlen, _, found := findDDLColumnConstraints(fs, "email")
		if !found {
			t.Fatal("expected found=true")
		}
		if vlen != 255 {
			t.Errorf("VarcharLen = %d, want 255", vlen)
		}
	})

	t.Run("CHECK enum extracted", func(t *testing.T) {
		_, enums, found := findDDLColumnConstraints(fs, "role")
		if !found {
			t.Fatal("expected found=true")
		}
		if len(enums) != 2 {
			t.Fatalf("expected 2 enums, got %d: %v", len(enums), enums)
		}
	})
}
