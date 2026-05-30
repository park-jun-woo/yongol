//ff:func feature=validate type=test control=selection
//ff:what TestStepIsPresent — step.isPresent 모든 kind 존재 여부 판정 분기 검증

package validate

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestStepIsPresent(t *testing.T) {
	t.Run("no kinds is present", func(t *testing.T) {
		s := step{Name: "x"}
		if !s.isPresent(&yongol.Fullstack{}) {
			t.Error("expected present for empty kinds")
		}
	})

	t.Run("one kind missing", func(t *testing.T) {
		s := step{Kinds: []yongol.SSOTKind{yongol.KindDDL, yongol.KindOpenAPI}}
		fs := &yongol.Fullstack{DDLTables: []ddl.Table{}}
		if s.isPresent(fs) {
			t.Error("expected not present when one kind missing")
		}
	})

	t.Run("all kinds present", func(t *testing.T) {
		s := step{Kinds: []yongol.SSOTKind{yongol.KindDDL}}
		fs := &yongol.Fullstack{DDLTables: []ddl.Table{}}
		if !s.isPresent(fs) {
			t.Error("expected present when all kinds present")
		}
	})
}
