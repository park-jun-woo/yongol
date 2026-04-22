//ff:func feature=rule type=test control=iteration dimension=1
//ff:what splitPascal — PascalCase 분해 (빈/단일/복합/전부-대문자/소문자 시작)

package rule

import (
	"reflect"
	"testing"
)

func TestSplitPascal(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{"", nil},
		{"Url", []string{"Url"}},
		{"OrgId", []string{"Org", "Id"}},
		{"URLBox", []string{"URL", "Box"}},
		{"ID", []string{"ID"}},
		{"Email", []string{"Email"}},
		{"lower", []string{"lower"}}, // leading lowercase fallback
	}
	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			got := splitPascal(c.in)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("splitPascal(%q) = %v; want %v", c.in, got, c.want)
			}
		})
	}
}
