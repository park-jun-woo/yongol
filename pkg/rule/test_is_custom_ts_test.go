//ff:func feature=rule type=test control=iteration dimension=1
//ff:what IsCustomTS — Ground.Flags["customTS.<name>"] 로 claim 조회

package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/pkg/toulmin"
)

func TestIsCustomTS(t *testing.T) {
	cases := []struct {
		name  string
		flags StringSet
		claim string
		want  bool
	}{
		{"exists", StringSet{"customTS.Foo": true}, "Foo", true},
		{"missing", StringSet{"customTS.Foo": true}, "Bar", false},
		{"empty-flags", StringSet{}, "Foo", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := toulmin.NewContext()
			ctx.Set("ground", &Ground{Flags: c.flags})
			ctx.Set("claim", c.claim)
			got, _ := IsCustomTS(ctx, nil)
			if got != c.want {
				t.Fatalf("IsCustomTS(flags=%v,claim=%q) = %v; want %v", c.flags, c.claim, got, c.want)
			}
		})
	}
}
