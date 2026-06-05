//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestSchemaFieldDecl — schemaFieldDecl required/optional Pydantic 필드 선언 검증
package fastapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestSchemaFieldDecl(t *testing.T) {
	cases := []struct {
		name  string
		field ir.BodyFieldMeta
		want  string
	}{
		{
			name:  "required str",
			field: ir.BodyFieldMeta{Name: "title", Format: "", Required: true},
			want:  "    title: str\n",
		},
		{
			name:  "required int",
			field: ir.BodyFieldMeta{Name: "count", Format: "int64", Required: true},
			want:  "    count: int\n",
		},
		{
			name:  "optional str",
			field: ir.BodyFieldMeta{Name: "note", Format: "email", Required: false},
			want:  "    note: Optional[str] = None\n",
		},
		{
			name:  "optional float",
			field: ir.BodyFieldMeta{Name: "ratio", Format: "double", Required: false},
			want:  "    ratio: Optional[float] = None\n",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := schemaFieldDecl(c.field); got != c.want {
				t.Errorf("schemaFieldDecl(%+v) = %q, want %q", c.field, got, c.want)
			}
		})
	}
}
