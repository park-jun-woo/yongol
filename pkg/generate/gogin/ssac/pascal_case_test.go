//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what pascalCase 단위 테스트

package ssac

import "testing"

func TestPascalCase(t *testing.T) {
	cases := map[string]string{
		"":         "",
		"id":       "Id",
		"org_id":   "OrgId",
		"ID":       "ID",
		"Workflow": "Workflow",
		"name":     "Name",
		"a_b_c":    "ABC",
	}
	for in, want := range cases {
		if got := pascalCase(in); got != want {
			t.Errorf("pascalCase(%q) = %q, want %q", in, got, want)
		}
	}
}
