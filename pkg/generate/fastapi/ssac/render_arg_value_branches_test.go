//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestRenderArgValueBranches — renderArgValue 미커버 분기(LocLiteral/LocVar no-source) 보강

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderArgValueBranches(t *testing.T) {
	tests := []struct {
		name string
		arg  ir.FieldArg
		want string
	}{
		{
			name: "LocLiteral",
			arg:  ir.FieldArg{Location: ir.LocLiteral, Literal: ""},
			want: `""`,
		},
		{
			// LocVar with srcCol resolved but no Source -> returns srcCol only.
			name: "LocVarSrcColNoSource",
			arg:  ir.FieldArg{Location: ir.LocVar, SourceColumn: "org_id"},
			want: "org_id",
		},
		{
			// LocVar with Source but empty srcCol falls back to col (ColumnName).
			name: "LocVarSrcColFallbackToCol",
			arg:  ir.FieldArg{Location: ir.LocVar, ColumnName: "id"},
			want: "id",
		},
		{
			// LocVar no source, no columns at all -> empty string.
			name: "LocVarEmpty",
			arg:  ir.FieldArg{Location: ir.LocVar},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := renderArgValue(tt.arg); got != tt.want {
				t.Errorf("renderArgValue() = %q, want %q", got, tt.want)
			}
		})
	}
}
