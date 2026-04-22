//ff:func feature=rule type=test control=sequence dimension=1
//ff:what stripTypePrefix — 슬라이스/포인터 prefix 제거

package ground

import "testing"

func TestStripTypePrefix_Cases(t *testing.T) {
	cases := []struct{ in, want string }{
		{"Workflow", "Workflow"},
		{"[]Workflow", "Workflow"},
		{"*User", "User"},
		{"", ""},
	}
	for _, c := range cases {
		if got := stripTypePrefix(c.in); got != c.want {
			t.Errorf("stripTypePrefix(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
