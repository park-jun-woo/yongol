//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what TestQ11SqlPackagePgxV5_AbsentFails — sql_package 누락 은 Q-11 에러

package query

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestQ11SqlPackagePgxV5_AbsentFails asserts that an absent sql_package
// fires Q-11 and the message explicitly reports the absence.
func TestQ11SqlPackagePgxV5_AbsentFails(t *testing.T) {
	dir := t.TempDir()
	writeSqlcYaml(t, dir, "")
	fs := &yongol.Fullstack{SpecsDir: dir}
	diags := q11SqlPackagePgxV5(fs)
	if len(diags) == 0 {
		t.Fatalf("absent sql_package must fire Q-11, got no diagnostics")
	}
	if !strings.Contains(diags[0].Message, "absent") {
		t.Errorf("message should note absence: %q", diags[0].Message)
	}
}
