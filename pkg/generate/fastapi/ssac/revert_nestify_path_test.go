//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestRevertNestifyPath — :param → {param} 경로 복원

package ssac

import "testing"

func TestRevertNestifyPath(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"/workflow/:id", "/workflow/{id}"},
		{"/workflow/:id/execute", "/workflow/{id}/execute"},
		{"/org/:org_id/item/:item_id", "/org/{org_id}/item/{item_id}"},
		{"/static/path", "/static/path"},
		{"", ""},
	}
	for _, c := range cases {
		if got := revertNestifyPath(c.in); got != c.want {
			t.Errorf("revertNestifyPath(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
