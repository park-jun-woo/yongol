//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what collectQueryNames — SQLcQueries의 name 세트 구성 검증

package manifest

import (
	"testing"

	sqlc "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCollectQueryNames(t *testing.T) {
	cases := []struct {
		name   string
		fs     *yongol.Fullstack
		wantN  int
		wantIn []string
	}{
		{name: "nil_queries", fs: &yongol.Fullstack{}, wantN: 0},
		{
			name: "collects_names",
			fs: &yongol.Fullstack{SQLcQueries: []sqlc.QuerySpec{
				{Name: "GetUser"}, {Name: "ListOrders"},
			}},
			wantN:  2,
			wantIn: []string{"GetUser", "ListOrders"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runBoolSetCase(t, collectQueryNames(c.fs), c.wantN, c.wantIn)
		})
	}
}
