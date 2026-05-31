//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what render_root_params_unit_test — findRootElement / renderUseParams 단위 테스트
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderUseParams(t *testing.T) {
	cases := []struct {
		name   string
		params []stmlparser.ParamBind
		want   string
	}{
		{"no params", nil, ""},
		{"no route params", []stmlparser.ParamBind{{Name: "Q", Source: "query.q"}}, ""},
		{"single route param", []stmlparser.ParamBind{{Name: "ID", Source: "route.ID"}}, "const { ID } = useParams()"},
		{
			"multiple route params",
			[]stmlparser.ParamBind{
				{Name: "ResID", Source: "route.ResID"},
				{Name: "Tab", Source: "route.Tab"},
			},
			"const { ResID, Tab } = useParams()",
		},
		{
			"dedup repeated route param",
			[]stmlparser.ParamBind{
				{Name: "ID", Source: "route.ID"},
				{Name: "ID2", Source: "route.ID"},
			},
			"const { ID } = useParams()",
		},
		{
			"mixed route and non-route",
			[]stmlparser.ParamBind{
				{Name: "ID", Source: "route.ID"},
				{Name: "Q", Source: "query.q"},
			},
			"const { ID } = useParams()",
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := renderUseParams(c.params); got != c.want {
				t.Errorf("renderUseParams = %q, want %q", got, c.want)
			}
		})
	}
}
