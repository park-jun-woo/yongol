//ff:func feature=rule type=test control=iteration dimension=1
//ff:what IsImplicitVar — Ground.Flags["implicit.<name>"] 로 claim 조회

package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/pkg/toulmin"
)

func TestIsImplicitVar(t *testing.T) {
	cases := []struct {
		name  string
		flags StringSet
		claim string
		want  bool
	}{
		{"implicit-user", StringSet{"implicit.user_id": true}, "user_id", true},
		{"not-implicit", StringSet{"implicit.user_id": true}, "tenant_id", false},
		{"empty", StringSet{}, "anything", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := toulmin.NewContext()
			ctx.Set("ground", &Ground{Flags: c.flags})
			ctx.Set("claim", c.claim)
			got, _ := IsImplicitVar(ctx, nil)
			if got != c.want {
				t.Fatalf("IsImplicitVar(flags=%v,claim=%q) = %v; want %v", c.flags, c.claim, got, c.want)
			}
		})
	}
}
