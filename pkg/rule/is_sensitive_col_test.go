//ff:func feature=rule type=test control=iteration dimension=1
//ff:what IsSensitiveCol — Ground.Flags["sensitive"] 여부만 확인

package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/pkg/toulmin"
)

func TestIsSensitiveCol(t *testing.T) {
	cases := []struct {
		name  string
		flags StringSet
		want  bool
	}{
		{"sensitive", StringSet{"sensitive": true}, true},
		{"not-sensitive", StringSet{}, false},
		{"nosensitive-only", StringSet{"nosensitive": true}, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := toulmin.NewContext()
			ctx.Set("ground", &Ground{Flags: c.flags})
			got, _ := IsSensitiveCol(ctx, nil)
			if got != c.want {
				t.Fatalf("IsSensitiveCol(%v) = %v; want %v", c.flags, got, c.want)
			}
		})
	}
}
