//ff:func feature=gen-hurl type=test control=iteration dimension=1
//ff:what canResolvePathParams 해소 가능 여부 테스트
package hurl

import "testing"

func TestCanResolvePathParams(t *testing.T) {
	type tc struct {
		name     string
		path     string
		captures map[string]bool
		want     bool
	}
	cases := []tc{
		{"no params", "/users/me", nil, true},
		{"direct match", "/workflows/{workflow_id}", map[string]bool{"workflow_id": true}, true},
		{"derived from plural", "/workflows/{id}", map[string]bool{"workflow_id": true}, true},
		{"hyphen resource derived", "/audit-logs/{id}", map[string]bool{"audit_log_id": true}, true},
		{"hyphen resource missing", "/audit-logs/{id}", nil, false},
		{"pascal param id", "/gigs/{ID}", map[string]bool{"gig_id": true}, true},
		{"multi param mixed", "/workflows/{id}/actions/{id}", map[string]bool{"workflow_id": true, "action_id": true}, true},
		{"first ok second missing", "/workflows/{id}/actions/{id}", map[string]bool{"workflow_id": true}, false},
	}
	for _, c := range cases {
		if got := canResolvePathParams(c.path, c.captures); got != c.want {
			t.Errorf("canResolvePathParams(%q) = %v, want %v (captures=%v)", c.path, got, c.want, c.captures)
		}
	}
}
