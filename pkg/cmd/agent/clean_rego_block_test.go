//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestCleanRegoBlock — 중복 package/import/default allow 라인 제거 검증

package agent

import (
	"strings"
	"testing"
)

func TestCleanRegoBlock(t *testing.T) {
	in := "package authz\nimport future.keywords\ndefault allow = false\nallow {\n  input.x == 1\n}"
	got := cleanRegoBlock(in)
	if strings.Contains(got, "package ") {
		t.Errorf("package line not removed: %q", got)
	}
	if strings.Contains(got, "import ") {
		t.Errorf("import line not removed: %q", got)
	}
	if strings.Contains(got, "default allow") {
		t.Errorf("default allow line not removed: %q", got)
	}
	if !strings.Contains(got, "input.x == 1") {
		t.Errorf("body lost: %q", got)
	}
}
