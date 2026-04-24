//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what TestQ11SqlPackagePgxV5_DatabaseSqlFails — database/sql 는 Q-11 에러

package query

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestQ11SqlPackagePgxV5_DatabaseSqlFails asserts that database/sql
// surfaces a Q-11 diagnostic with the current value echoed back.
func TestQ11SqlPackagePgxV5_DatabaseSqlFails(t *testing.T) {
	dir := t.TempDir()
	writeSqlcYaml(t, dir, "database/sql")
	fs := &yongol.Fullstack{SpecsDir: dir}
	diags := q11SqlPackagePgxV5(fs)
	if len(diags) == 0 {
		t.Fatalf("database/sql must fire Q-11, got no diagnostics")
	}
	if !strings.Contains(diags[0].Message, "[Q-11]") {
		t.Errorf("rule id missing: %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, `"database/sql"`) {
		t.Errorf("current value missing from message: %q", diags[0].Message)
	}
}
