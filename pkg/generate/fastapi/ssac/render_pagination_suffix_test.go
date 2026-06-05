//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestRenderPaginationSuffix — renderPaginationSuffix limit/offset SQLAlchemy 접미사·미매칭 무시 검증
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderPaginationSuffix(t *testing.T) {
	cases := []struct {
		name    string
		pagArgs []ir.FieldArg
		want    string
	}{
		{
			name: "limit and offset",
			pagArgs: []ir.FieldArg{
				{ColumnName: "limit", Location: ir.LocQuery},
				{ColumnName: "offset", Location: ir.LocQuery},
			},
			want: ".limit(limit).offset(offset)",
		},
		{
			name: "per_page and page_offset",
			pagArgs: []ir.FieldArg{
				{ColumnName: "per_page", Location: ir.LocQuery},
				{ColumnName: "page_offset", Location: ir.LocQuery},
			},
			want: ".limit(per_page).offset(page_offset)",
		},
		{
			name:    "unmatched key ignored",
			pagArgs: []ir.FieldArg{{ColumnName: "cursor", Location: ir.LocQuery}},
			want:    "",
		},
		{
			name:    "empty",
			pagArgs: nil,
			want:    "",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := renderPaginationSuffix(c.pagArgs); got != c.want {
				t.Errorf("renderPaginationSuffix() = %q, want %q", got, c.want)
			}
		})
	}
}
