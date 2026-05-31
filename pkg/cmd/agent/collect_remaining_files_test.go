//ff:func feature=agent type=test control=sequence
//ff:what TestCollectRemainingFiles — ERROR만 file:line 요약으로 정렬·중복제거하는지 검증
package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestCollectRemainingFiles(t *testing.T) {
	diags := []diagnostic.Diagnostic{
		{Level: diagnostic.LevelError, File: "/specs/db/z.sql", Line: 2, Message: "zebra"},
		{Level: diagnostic.LevelWarning, File: "/specs/db/a.sql", Line: 1, Message: "warn"}, // skipped
		{Level: diagnostic.LevelError, File: "/specs/api/a.yaml", Line: 1, Message: "alpha"},
		{Level: diagnostic.LevelError, File: "/specs/db/z.sql", Line: 2, Message: "zebra"}, // dup
	}
	got := collectRemainingFiles(diags, "/specs")
	if len(got) != 2 {
		t.Fatalf("got %d entries, want 2: %v", len(got), got)
	}
	// Sorted: "api/a.yaml:1 alpha" < "db/z.sql:2 zebra".
	if got[0] != "api/a.yaml:1 alpha" {
		t.Errorf("got[0] = %q", got[0])
	}
	if got[1] != "db/z.sql:2 zebra" {
		t.Errorf("got[1] = %q", got[1])
	}
}
