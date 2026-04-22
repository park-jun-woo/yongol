//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestDF_06_DBInit_HasDeferClose — blockDBInit 템플릿이 defer conn.Close() 를 포함하는지 회귀 방지

package boot

import (
	"strings"
	"testing"
)

// TestDF_06_DBInit_HasDeferClose asserts the db-init template keeps
// `defer conn.Close()` (Phase003 DF-06). Also asserts the sql.Open err
// check immediately follows (DF guard). If either guard is removed a
// panic/leak regression would ship to every generated backend.
func TestDF_06_DBInit_HasDeferClose(t *testing.T) {
	block := blockDBInit(nil, "example.com/zenflow")
	lines := strings.Join(block.Lines, "\n")
	if !strings.Contains(lines, "defer conn.Close()") {
		t.Fatalf("db-init must defer conn.Close() (DF-06), got:\n%s", lines)
	}
	if !strings.Contains(lines, `conn, err := sql.Open("postgres"`) {
		t.Fatalf("db-init must assign sql.Open error, got:\n%s", lines)
	}
	if !strings.Contains(lines, "if err != nil {") {
		t.Fatalf("db-init must guard sql.Open error (DF-01 family), got:\n%s", lines)
	}
}
