//ff:func feature=rule type=test control=iteration dimension=1
//ff:what IsArchived — Ground.Flags["archived"] 여부만 확인

package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/pkg/toulmin"
)

func TestIsArchived(t *testing.T) {
	cases := []struct {
		name  string
		flags StringSet
		want  bool
	}{
		{"archived", StringSet{"archived": true}, true},
		{"not-archived", StringSet{}, false},
		{"unrelated-flag", StringSet{"dto": true}, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := toulmin.NewContext()
			ctx.Set("ground", &Ground{Flags: c.flags})
			got, _ := IsArchived(ctx, nil)
			if got != c.want {
				t.Fatalf("IsArchived(%v) = %v; want %v", c.flags, got, c.want)
			}
		})
	}
}
