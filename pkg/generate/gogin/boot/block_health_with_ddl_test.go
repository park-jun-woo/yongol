//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockHealth — /health (liveness) + /ready (readiness) 등록
package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockHealth_WithDDL(t *testing.T) {
	fs := &yongol.Fullstack{DDLTables: []ddl.Table{{Columns: map[string]ddl.Column{"id": {}}}}}
	block := blockHealth(fs)
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, `r.GET("/ready", readyHandlerWithDB(pool))`) {
		t.Errorf("DDL /ready should delegate to readyHandlerWithDB, got:\n%s", body)
	}
	if len(block.Funcs) != 1 {
		t.Errorf("DDL health should emit readyHandlerWithDB helper, got %d funcs", len(block.Funcs))
	}
	if !strings.Contains(strings.Join(block.Imports, "\n"), "pgxpool") {
		t.Errorf("DDL health must import pgxpool, got:\n%v", block.Imports)
	}
}
