//ff:func feature=gen-hurl type=test control=iteration dimension=1
//ff:what snakeHurlName 정규화 테스트
package hurl

import "testing"

func TestSnakeHurlName(t *testing.T) {
	cases := map[string]string{
		"audit-logs":    "audit_logs",
		"AuditLog":      "audit_log",
		"GigID":         "gig_id",
		"id":            "id",
		"execution-log": "execution_log",
		"":              "",
		"UserID":        "user_id",
		"snake_already": "snake_already",
	}
	for in, want := range cases {
		if got := snakeHurlName(in); got != want {
			t.Errorf("snakeHurlName(%q) = %q, want %q", in, got, want)
		}
	}
}
