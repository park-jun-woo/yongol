//ff:func feature=generate type=test control=sequence
//ff:what TestZeroCov — 0% util 함수 (isCopiedExtension / isYongolManaged / mergeFieldlessOps / ResolveBackendType / WithMigration / appendChildNodeFormActions) 회귀
package generate

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestLogMigrationWarnings_ZeroCov(t *testing.T) {
	// nil logger → no panic, early return.
	logMigrationWarnings(MigrationHook{Logger: nil}, []diagnostic.Diagnostic{{Level: diagnostic.LevelWarning, Message: "w"}})

	var buf bytes.Buffer
	diags := []diagnostic.Diagnostic{
		{Level: diagnostic.LevelWarning, Message: "warn1"},
		{Level: diagnostic.LevelError, Message: "err1"},
	}
	logMigrationWarnings(MigrationHook{Logger: &buf}, diags)
	s := buf.String()
	if !strings.Contains(s, "warn1") {
		t.Errorf("missing warning: %q", s)
	}
	if strings.Contains(s, "err1") {
		t.Errorf("error should not be logged: %q", s)
	}
}
