//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestResolveDataKey — ColumnName/Key/Field 우선순위 키 추출
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestResolveDataKey(t *testing.T) {
	cases := []struct {
		name string
		arg  ir.FieldArg
		want string
	}{
		{"ColumnNamePreferred", ir.FieldArg{ColumnName: "org_id", Key: "OrgID", Field: ".X"}, "org_id"},
		{"KeyFallback", ir.FieldArg{Key: "OrgID"}, "org_id"},
		{"FieldFallback", ir.FieldArg{Field: ".CreatedAt"}, "created_at"},
		{"FieldNoDot", ir.FieldArg{Field: "Status"}, "status"},
		{"AllEmpty", ir.FieldArg{}, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := resolveDataKey(c.arg); got != c.want {
				t.Errorf("resolveDataKey() = %q, want %q", got, c.want)
			}
		})
	}
}
