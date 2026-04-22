//ff:func feature=rule type=test control=sequence dimension=1
//ff:what refName — $ref 경로의 마지막 segment 반환

package ground

import "testing"

func TestRefName_Cases(t *testing.T) {
	cases := []struct{ in, want string }{
		{"#/components/schemas/Workflow", "Workflow"},
		{"#/components/schemas/User", "User"},
		{"X", "X"}, // no slash → return as is
		{"", ""},
	}
	for _, c := range cases {
		if got := refName(c.in); got != c.want {
			t.Errorf("refName(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
