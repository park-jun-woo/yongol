//ff:func feature=generate type=test control=sequence
//ff:what TestZeroCov — 0% util 함수 (isCopiedExtension / isYongolManaged / mergeFieldlessOps / ResolveBackendType / WithMigration / appendChildNodeFormActions) 회귀
package generate

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestLogIncrementalMigration_ZeroCov(t *testing.T) {
	var buf bytes.Buffer
	res := &migration.Result{
		MigrationFile: "0002_x.up.sql",
		OpsCount:      1,
		Operations:    []migration.Operation{migration.DropCheck{Table: "t", Name: "c"}},
	}
	logIncrementalMigration(MigrationHook{Logger: &buf}, res)
	s := buf.String()
	if !strings.Contains(s, "incremental") || !strings.Contains(s, "0002_x.up.sql") || !strings.Contains(s, "drop check c") {
		t.Errorf("log output = %q", s)
	}
}
