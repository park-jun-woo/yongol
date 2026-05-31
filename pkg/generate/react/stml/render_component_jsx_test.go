//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what render_component_static_unit_test — renderComponentJSX / renderStaticActionJSX 단위 테스트
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderComponentJSX(t *testing.T) {
	cases := []struct {
		name    string
		comp    stmlparser.ComponentRef
		dataVar string
		indent  int
		want    string
	}{
		{
			"no bind",
			stmlparser.ComponentRef{Name: "DatePicker"},
			"data", 2,
			"  <DatePicker />",
		},
		{
			"simple bind",
			stmlparser.ComponentRef{Name: "Avatar", Bind: "url"},
			"data", 0,
			"<Avatar data={data.url} />",
		},
		{
			"nested bind uses optional chaining",
			stmlparser.ComponentRef{Name: "Card", Bind: "user.profile.name"},
			"resp", 4,
			"    <Card data={resp.user?.profile?.name} />",
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := renderComponentJSX(c.comp, c.dataVar, c.indent); got != c.want {
				t.Errorf("renderComponentJSX = %q, want %q", got, c.want)
			}
		})
	}
}
