//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockLoggerInit — slog 기본 핸들러 초기화 (JSON/Text, LOG_LEVEL, redact)

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockLoggerInit(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{{Columns: map[string]ddl.Column{"ssn": {Sensitive: true}}}},
	}
	block := blockLoggerInit(fs)
	if block.Name != "logger-init" {
		t.Errorf("name = %q, want logger-init", block.Name)
	}
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, `parseLogLevel(os.Getenv("LOG_LEVEL"))`) {
		t.Errorf("must read LOG_LEVEL, got:\n%s", body)
	}
	if !strings.Contains(body, `buildSensitiveKeys([]string{"ssn"})`) {
		t.Errorf("must inject @sensitive columns, got:\n%s", body)
	}
	if !strings.Contains(body, "slog.SetDefault(") {
		t.Errorf("must set default slog, got:\n%s", body)
	}
	imp := strings.Join(block.Imports, "\n")
	if !strings.Contains(imp, `"github.com/park-jun-woo/ssac/pkg/redact"`) {
		t.Errorf("must import redact, got:\n%s", imp)
	}
	// Three helper funcs emitted at package scope.
	if len(block.Funcs) != 3 {
		t.Errorf("expected 3 helper funcs, got %d", len(block.Funcs))
	}
}
