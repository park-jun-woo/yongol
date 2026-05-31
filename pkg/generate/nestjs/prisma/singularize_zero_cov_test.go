//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestIsPrimaryKey_ZeroCov — isPrimaryKey 포함/미포함 분기 검증
package prisma

import (
	"testing"
)

func TestSingularize_ZeroCov(t *testing.T) {
	cases := map[string]string{
		"categories": "category",
		"classes":    "class",
		"boxes":      "box",
		"users":      "user",
		"address":    "address", // ss → unchanged by last branch
		"item":       "item",
	}
	for in, want := range cases {
		if got := singularize(in); got != want {
			t.Errorf("singularize(%q)=%q want %q", in, got, want)
		}
	}
}
