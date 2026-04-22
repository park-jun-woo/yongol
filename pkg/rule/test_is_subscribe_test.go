//ff:func feature=rule type=test control=iteration dimension=1
//ff:what IsSubscribe — Ground.Flags["subscribe"] 여부만 확인

package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/pkg/toulmin"
)

func TestIsSubscribe(t *testing.T) {
	cases := []struct {
		name  string
		flags StringSet
		want  bool
	}{
		{"subscribe", StringSet{"subscribe": true}, true},
		{"not-subscribe", StringSet{}, false},
		{"unrelated", StringSet{"dto": true}, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := toulmin.NewContext()
			ctx.Set("ground", &Ground{Flags: c.flags})
			got, _ := IsSubscribe(ctx, nil)
			if got != c.want {
				t.Fatalf("IsSubscribe(%v) = %v; want %v", c.flags, got, c.want)
			}
		})
	}
}
