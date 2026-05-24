//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what diagnoseSqlcPackageEntry — pgx/v5 pass, 빈 값, 다른 값 3분기 진단 검증

package query

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestDiagnoseSqlcPackageEntry(t *testing.T) {
	t.Run("pgx/v5 returns nil", func(t *testing.T) {
		d := diagnoseSqlcPackageEntry(0, "pgx/v5")
		if d != nil {
			t.Fatalf("expected nil, got %+v", d)
		}
	})

	t.Run("empty string reports absent", func(t *testing.T) {
		d := diagnoseSqlcPackageEntry(2, "")
		if d == nil {
			t.Fatal("expected diagnostic, got nil")
		}
		if d.File != "db/sqlc.yaml" {
			t.Errorf("expected File=db/sqlc.yaml, got %s", d.File)
		}
		if d.Phase != diagnostic.PhaseValidate {
			t.Errorf("expected PhaseValidate, got %v", d.Phase)
		}
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %v", d.Level)
		}
		if !strings.Contains(d.Message, "[Q-11]") {
			t.Errorf("expected Q-11 in message, got %s", d.Message)
		}
		if !strings.Contains(d.Message, "sql[2]") {
			t.Errorf("expected sql[2] in message, got %s", d.Message)
		}
		if !strings.Contains(d.Message, "(absent; sqlc defaults to database/sql)") {
			t.Errorf("expected absent fallback in message, got %s", d.Message)
		}
		if !strings.Contains(d.Advice, "pgx/v5") {
			t.Errorf("expected pgx/v5 in advice, got %s", d.Advice)
		}
	})

	t.Run("pgx/v4 reports quoted value", func(t *testing.T) {
		d := diagnoseSqlcPackageEntry(1, "pgx/v4")
		if d == nil {
			t.Fatal("expected diagnostic, got nil")
		}
		if !strings.Contains(d.Message, "[Q-11]") {
			t.Errorf("expected Q-11 in message, got %s", d.Message)
		}
		if !strings.Contains(d.Message, "sql[1]") {
			t.Errorf("expected sql[1] in message, got %s", d.Message)
		}
		if !strings.Contains(d.Message, `"pgx/v4"`) {
			t.Errorf("expected quoted pgx/v4 in message, got %s", d.Message)
		}
	})

	t.Run("database/sql reports quoted value", func(t *testing.T) {
		d := diagnoseSqlcPackageEntry(0, "database/sql")
		if d == nil {
			t.Fatal("expected diagnostic, got nil")
		}
		if !strings.Contains(d.Message, `"database/sql"`) {
			t.Errorf("expected quoted database/sql in message, got %s", d.Message)
		}
	})
}
