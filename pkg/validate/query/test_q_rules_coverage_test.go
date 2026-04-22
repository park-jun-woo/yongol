//ff:func feature=validate type=test control=iteration dimension=2 topic=query-structural
//ff:what Q-01 / Q-04 / Q-07 / Q-08 positive 테스트 — 규칙 ID별 최소 1 발화 보장

package query

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
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

// TestQ04ManyLimitFires covers Q-04: :many query missing LIMIT fires.
func TestQ04ManyLimitFires(t *testing.T) {
	fs := &yongol.Fullstack{SQLcQueries: loadSpecs(t, "q04_many_no_limit.sql")}
	diags := q04ManyLimit(fs)
	if len(diags) == 0 {
		t.Fatalf("want Q-04 diagnostic, got none")
	}
	if !strings.Contains(diags[0].Message, "[Q-04]") {
		t.Errorf("rule id missing: %q", diags[0].Message)
	}
}

// TestQ07SelectStarSensitiveFires covers Q-07: SELECT * on a table with a
// @sensitive column emits a WARNING.
func TestQ07SelectStarSensitiveFires(t *testing.T) {
	specs := loadSpecs(t, "q07_select_star.sql")
	fs := &yongol.Fullstack{
		SQLcQueries: specs,
		DDLTables: []ddl.Table{{
			Name:             "users",
			SensitiveColumns: map[string]bool{"password_hash": true},
		}},
	}
	diags := q07SelectStarSensitive(fs)
	if len(diags) == 0 {
		t.Fatalf("want Q-07 diagnostic, got none")
	}
	if !strings.Contains(diags[0].Message, "[Q-07]") {
		t.Errorf("rule id missing: %q", diags[0].Message)
	}
}

// TestQ08UnusedParamFires covers Q-08: declared param missing from body → ERROR.
func TestQ08UnusedParamFires(t *testing.T) {
	specs := loadSpecs(t, "q08_unused_param.sql")
	// Inject a synthetic unused parameter. The parser only records params
	// actually present in the SQL, so we augment Params here to exercise Q-08.
	for i := range specs {
		specs[i].Params = append(specs[i].Params, "GhostParam")
	}
	fs := &yongol.Fullstack{SQLcQueries: specs}
	diags := q08UnusedParam(fs)
	if len(diags) == 0 {
		t.Fatalf("want Q-08 diagnostic, got none")
	}
	if !strings.Contains(diags[0].Message, "[Q-08]") {
		t.Errorf("rule id missing: %q", diags[0].Message)
	}
}
