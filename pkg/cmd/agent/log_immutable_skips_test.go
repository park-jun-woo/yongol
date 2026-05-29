//ff:func feature=agent type=test control=iteration dimension=2
//ff:what TestLogImmutableSkips — immutable ERROR 파일만 1회 요약 출력, 중복/비-immutable/warning 제외 검증

package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestLogImmutableSkips(t *testing.T) {
	diags := []diagnostic.Diagnostic{
		{Level: diagnostic.LevelError, File: "/specs/tests/login.hurl", Message: "e1"},
		{Level: diagnostic.LevelError, File: "/specs/tests/login.hurl", Message: "e2"},  // same file, dedup
		{Level: diagnostic.LevelError, File: "/specs/db/users.sql", Message: "mutable"}, // not immutable
		{Level: diagnostic.LevelWarning, File: "/specs/features.yaml", Message: "warn"}, // warning skipped
		{Level: diagnostic.LevelError, File: "/specs/features.yaml", Message: "e3"},
	}
	var b strings.Builder
	logImmutableSkips(&b, diags, "/specs")
	out := b.String()

	if strings.Count(out, "skipped:") != 2 {
		t.Errorf("skipped lines = %d, want 2\n%s", strings.Count(out, "skipped:"), out)
	}
	if !strings.Contains(out, "tests/login.hurl") || !strings.Contains(out, "features.yaml") {
		t.Errorf("output missing immutable files: %s", out)
	}
	if strings.Contains(out, "users.sql") {
		t.Errorf("mutable file should not be reported: %s", out)
	}
}
