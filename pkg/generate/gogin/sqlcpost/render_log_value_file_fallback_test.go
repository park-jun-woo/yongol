//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestRenderLogValueFile_EmptyColumnOrder — ColumnOrder 비면 정렬된 컬럼키로 fallback 검증
package sqlcpost

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestRenderLogValueFile_EmptyColumnOrder(t *testing.T) {
	table := ddl.Table{
		Name: "accounts",
		Columns: map[string]ddl.Column{
			"id":     {Name: "id", RawType: "UUID"},
			"secret": {Name: "secret", RawType: "TEXT", Sensitive: true},
		},
		// ColumnOrder intentionally empty -> sortedColumnKeys fallback.
	}
	src, err := renderLogValueFile(table)
	if err != nil {
		t.Fatalf("renderLogValueFile: %v", err)
	}
	if !strings.Contains(src, "LogValue() slog.Value") {
		t.Errorf("expected LogValue method in output:\n%s", src)
	}
	if !strings.Contains(src, "[REDACTED]") {
		t.Errorf("expected sensitive column to be redacted:\n%s", src)
	}
}
