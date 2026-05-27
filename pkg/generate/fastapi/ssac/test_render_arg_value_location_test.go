//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestRenderArgValueLocation — FieldArg.Location 기반 Python 식 생성 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderArgValueLocation(t *testing.T) {
	tests := []struct {
		name string
		arg  ir.FieldArg
		want string
	}{
		{
			name: "LocPath",
			arg:  ir.FieldArg{Location: ir.LocPath, ColumnName: "org_id", Source: "request", Field: ".OrgID"},
			want: "org_id",
		},
		{
			name: "LocBody",
			arg:  ir.FieldArg{Location: ir.LocBody, ColumnName: "title", Source: "request", Field: ".Title"},
			want: "body.title",
		},
		{
			name: "LocQuery",
			arg:  ir.FieldArg{Location: ir.LocQuery, ColumnName: "per_page", Source: "request", Field: ".PerPage"},
			want: "per_page",
		},
		{
			name: "LocUser",
			arg:  ir.FieldArg{Location: ir.LocUser, ColumnName: "org_id", Source: "currentUser", Field: ".OrgID"},
			want: `current_user["org_id"]`,
		},
		{
			name: "LocVarWithField",
			arg:  ir.FieldArg{Location: ir.LocVar, ColumnName: "id", Source: "wf"},
			want: "wf.id",
		},
		{
			name: "LocVarNoField",
			arg:  ir.FieldArg{Location: ir.LocVar, Source: "wf"},
			want: "wf",
		},
		{
			name: "QuotedLiteral",
			arg:  ir.FieldArg{Literal: "active", IsQuoted: true},
			want: `"active"`,
		},
		{
			name: "UnquotedLiteral",
			arg:  ir.FieldArg{Literal: "42"},
			want: "42",
		},
		{
			name: "FallbackNoLocation",
			arg:  ir.FieldArg{Source: "request", Field: ".ID", ColumnName: "id"},
			want: "id",
		},
		{
			name: "FallbackCurrentUser",
			arg:  ir.FieldArg{Source: "currentUser", Field: ".OrgID", ColumnName: "org_id"},
			want: `current_user["org_id"]`,
		},
		{
			name: "FallbackVariable",
			arg:  ir.FieldArg{Source: "wf", Field: ".Status", ColumnName: "status"},
			want: "wf.status",
		},
		{
			name: "FallbackNoFieldRequest",
			arg:  ir.FieldArg{Source: "request"},
			want: "params",
		},
		{
			name: "FallbackNoFieldCurrentUser",
			arg:  ir.FieldArg{Source: "currentUser"},
			want: "current_user",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := renderArgValue(tt.arg)
			if got != tt.want {
				t.Errorf("renderArgValue() = %q, want %q", got, tt.want)
			}
		})
	}
}
