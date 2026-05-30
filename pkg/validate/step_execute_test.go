//ff:func feature=validate type=test control=selection
//ff:what TestStepExecute — step.execute missing/skip/run 정책 분기 검증

package validate

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestStepExecute(t *testing.T) {
	t.Run("missing SSOT", func(t *testing.T) {
		s := step{Name: "ddl", Kinds: []yongol.SSOTKind{yongol.KindDDL}}
		res := s.execute(&yongol.Fullstack{}, &config{})
		if res.Status != StatusMissing {
			t.Errorf("expected StatusMissing, got %v", res.Status)
		}
	})

	t.Run("present but not wired", func(t *testing.T) {
		s := step{Name: "ddl", Kinds: []yongol.SSOTKind{yongol.KindDDL}, Run: nil}
		fs := &yongol.Fullstack{DDLTables: []ddl.Table{}}
		res := s.execute(fs, &config{})
		if res.Status != StatusSkip {
			t.Errorf("expected StatusSkip, got %v", res.Status)
		}
	})

	t.Run("run returns pass", func(t *testing.T) {
		s := step{
			Name:  "ddl",
			Kinds: []yongol.SSOTKind{yongol.KindDDL},
			Run:   func(*yongol.Fullstack) []diagnostic.Diagnostic { return nil },
		}
		fs := &yongol.Fullstack{DDLTables: []ddl.Table{}}
		res := s.execute(fs, &config{})
		if res.Status != StatusPass {
			t.Errorf("expected StatusPass, got %v", res.Status)
		}
	})

	t.Run("run returns failure", func(t *testing.T) {
		s := step{
			Name:  "ddl",
			Kinds: []yongol.SSOTKind{yongol.KindDDL},
			Run: func(*yongol.Fullstack) []diagnostic.Diagnostic {
				return []diagnostic.Diagnostic{{Level: diagnostic.LevelError}}
			},
		}
		fs := &yongol.Fullstack{DDLTables: []ddl.Table{}}
		res := s.execute(fs, &config{})
		if res.Status != StatusFail {
			t.Errorf("expected StatusFail, got %v", res.Status)
		}
	})
}
