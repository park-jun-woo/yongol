//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-ddl
//ff:what buildReferencedTables — @model/@result 참조 DDL 테이블 수집 검증

package ssac_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestBuildReferencedTables(t *testing.T) {
	cases := []struct {
		name  string
		funcs []ssac.ServiceFunc
		want  map[string]bool
	}{
		{
			name:  "empty funcs",
			funcs: nil,
			want:  map[string]bool{},
		},
		{
			name: "model reference",
			funcs: []ssac.ServiceFunc{
				{Sequences: []ssac.Sequence{{Model: "Course.FindByID"}}},
			},
			want: map[string]bool{"course": true},
		},
		{
			name: "result reference",
			funcs: []ssac.ServiceFunc{
				{Sequences: []ssac.Sequence{{Result: &ssac.Result{Type: "Reservation"}}}},
			},
			want: map[string]bool{"reservation": true},
		},
		{
			name: "package sequence skipped",
			funcs: []ssac.ServiceFunc{
				{Sequences: []ssac.Sequence{{Package: "session", Model: "Session.Get"}}},
			},
			want: map[string]bool{},
		},
		{
			name: "primitive result skipped",
			funcs: []ssac.ServiceFunc{
				{Sequences: []ssac.Sequence{{Result: &ssac.Result{Type: "string"}}}},
			},
			want: map[string]bool{},
		},
		{
			name: "multiple funcs merged",
			funcs: []ssac.ServiceFunc{
				{Sequences: []ssac.Sequence{{Model: "User.Create"}}},
				{Sequences: []ssac.Sequence{{Model: "Course.FindByID"}}},
			},
			want: map[string]bool{"user": true, "course": true},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertReferencedTables(t, c.funcs, c.want)
		})
	}
}
