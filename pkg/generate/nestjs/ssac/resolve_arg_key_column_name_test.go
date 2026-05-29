//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestResolveArgKeyColumnName — ColumnName 우선 사용 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestResolveArgKeyColumnName(t *testing.T) {
	tests := []struct {
		name string
		arg  ir.FieldArg
		want string
	}{
		{
			name: "PreferColumnName",
			arg:  ir.FieldArg{ColumnName: "org_id", Key: "OrgID", Field: ".OrgID"},
			want: "org_id",
		},
		{
			name: "FallbackToKey",
			arg:  ir.FieldArg{Key: "OrgID"},
			want: "org_id",
		},
		{
			name: "FallbackToField",
			arg:  ir.FieldArg{Field: ".OrgID"},
			want: "org_id",
		},
		{
			name: "EmptyFallbackToID",
			arg:  ir.FieldArg{},
			want: "id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resolveArgKey(tt.arg)
			if got != tt.want {
				t.Errorf("resolveArgKey() = %q, want %q", got, tt.want)
			}
		})
	}
}
