//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what collectDDLNames — DDLTables의 name 세트 구성 검증

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCollectDDLNames(t *testing.T) {
	cases := []struct {
		name   string
		fs     *yongol.Fullstack
		wantN  int
		wantIn []string
	}{
		{name: "nil_tables", fs: &yongol.Fullstack{}, wantN: 0},
		{
			name:   "collects_names",
			fs:     &yongol.Fullstack{DDLTables: []ddl.Table{{Name: "users"}, {Name: "orders"}}},
			wantN:  2,
			wantIn: []string{"users", "orders"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runBoolSetCase(t, collectDDLNames(c.fs), c.wantN, c.wantIn)
		})
	}
}
