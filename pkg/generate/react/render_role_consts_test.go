//ff:func feature=gen-react type=test control=sequence
//ff:what renderRoleConsts — ROLES_* 상수 선언 방출/원본 role 값 보존(비식별자 살균은 이름만)/빈 목록 검증

package react

import (
	"strings"
	"testing"
)

func TestRenderRoleConsts(t *testing.T) {
	var sb strings.Builder
	renderRoleConsts(&sb, [][]string{
		{"admin", "manager"},
		{"super-admin"}, // sanitized in the name, original in the value
	})
	out := sb.String()

	if !strings.HasPrefix(out, "\n") {
		t.Errorf("output must start with a blank line, got %q", out)
	}
	assertContains(t, out, "const ROLES_admin_manager = ['admin', 'manager']\n")
	assertContains(t, out, "const ROLES_super_admin = ['super-admin']\n")

	// no role sets → just the leading newline, no const lines
	var empty strings.Builder
	renderRoleConsts(&empty, nil)
	if got := empty.String(); got != "\n" {
		t.Errorf("empty roleSets: got %q, want %q", got, "\n")
	}
}
