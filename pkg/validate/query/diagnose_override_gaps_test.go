//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what diagnoseOverrideGaps — 4분기 조합 (both/none/notNull only/nullable only) 진단 검증

package query

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestDiagnoseOverrideGaps(t *testing.T) {
	rule := pgtypeOverrideRule{
		RuleID:    "Q-99",
		DBType:    "uuid",
		PgPackage: "pgtype",
		PgType:    "UUID",
		Advice:    "Add uuid override",
	}

	t.Run("both present returns nil", func(t *testing.T) {
		diags := diagnoseOverrideGaps(rule, true, true)
		if diags != nil {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("both missing lists both sides", func(t *testing.T) {
		diags := diagnoseOverrideGaps(rule, false, false)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diagnostic, got %d", len(diags))
		}
		d := diags[0]
		if !strings.Contains(d.Message, "nullable=false and nullable=true") {
			t.Errorf("expected both sides in message, got %s", d.Message)
		}
		if !strings.Contains(d.Message, "Q-99") {
			t.Errorf("expected Q-99 in message, got %s", d.Message)
		}
		if !strings.Contains(d.Message, "uuid") {
			t.Errorf("expected db type uuid in message, got %s", d.Message)
		}
		if !strings.Contains(d.Message, "pgtype.UUID") {
			t.Errorf("expected pgtype.UUID in message, got %s", d.Message)
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
		if d.Advice != "Add uuid override" {
			t.Errorf("expected Advice='Add uuid override', got %s", d.Advice)
		}
	})

	t.Run("only nullable present missing notNull", func(t *testing.T) {
		diags := diagnoseOverrideGaps(rule, false, true)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diagnostic, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "nullable=false") {
			t.Errorf("expected nullable=false in message, got %s", diags[0].Message)
		}
		if strings.Contains(diags[0].Message, "nullable=true") {
			t.Errorf("should not mention nullable=true, got %s", diags[0].Message)
		}
	})

	t.Run("only notNull present missing nullable", func(t *testing.T) {
		diags := diagnoseOverrideGaps(rule, true, false)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diagnostic, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "nullable=true") {
			t.Errorf("expected nullable=true in message, got %s", diags[0].Message)
		}
		if strings.Contains(diags[0].Message, "nullable=false") {
			t.Errorf("should not mention nullable=false, got %s", diags[0].Message)
		}
	})
}
