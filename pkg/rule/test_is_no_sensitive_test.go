//ff:func feature=rule type=test control=iteration dimension=1
//ff:what IsNoSensitive — Ground.Flags["nosensitive"] 여부만 확인

package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/pkg/toulmin"
)

func TestIsNoSensitive(t *testing.T) {
	cases := []struct {
		name  string
		flags StringSet
		want  bool
	}{
		{"marked", StringSet{"nosensitive": true}, true},
		{"not-marked", StringSet{}, false},
		{"unrelated", StringSet{"sensitive": true}, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := toulmin.NewContext()
			ctx.Set("ground", &Ground{Flags: c.flags})
			got, _ := IsNoSensitive(ctx, nil)
			if got != c.want {
				t.Fatalf("IsNoSensitive(%v) = %v; want %v", c.flags, got, c.want)
			}
		})
	}
}
