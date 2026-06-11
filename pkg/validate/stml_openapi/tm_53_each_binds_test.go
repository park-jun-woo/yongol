//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what tm53EachBinds — data-each 항목 bind의 타입 대조(발화/통과/미상 스킵) 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM53EachBinds(t *testing.T) {
	itemFields := map[string]string{"url": "string", "size": "integer"}
	cases := []struct {
		name string
		bind stml.FieldBind
		want int
	}{
		{"img on string url silent", stml.FieldBind{Name: "url", Tag: "img"}, 0},
		{"img on integer fires", stml.FieldBind{Name: "size", Tag: "img"}, 1},
		{"scalar text silent", stml.FieldBind{Name: "url", Tag: "span"}, 0},
		{"unknown item skipped", stml.FieldBind{Name: "ghost", Tag: "span"}, 0},
	}
	for _, c := range cases {
		e := stml.EachBlock{Field: "photos", Binds: []stml.FieldBind{c.bind}}
		if d := tm53EachBinds(e, "GetThing", "p.html", itemFields); len(d) != c.want {
			t.Errorf("%s: want %d, got %+v", c.name, c.want, d)
		}
	}
}
