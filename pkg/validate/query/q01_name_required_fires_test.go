//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what Q-01 positive — `-- name:` 없는 SQL 은 ERROR

package query

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestQ01NameRequiredFires covers Q-01: SQL statement without `-- name:` fires.
func TestQ01NameRequiredFires(t *testing.T) {
	abs, err := filepath.Abs(filepath.Join("testdata", "q01_missing_name.sql"))
	if err != nil {
		t.Fatalf("abs: %v", err)
	}
	// q01 scans files referenced by QuerySpecs. We fake a QuerySpec that points
	// at the naked-SQL file so q01ScanForMissingName is invoked.
	fs := &yongol.Fullstack{SQLcQueries: []sqlc.QuerySpec{{
		Name: "dummy", File: abs, Line: 1,
	}}}
	diags := q01NameRequired(fs)
	if len(diags) == 0 {
		t.Fatalf("want Q-01 diagnostic, got none")
	}
	if !strings.Contains(diags[0].Message, "[Q-01]") {
		t.Errorf("rule id missing: %q", diags[0].Message)
	}
}
