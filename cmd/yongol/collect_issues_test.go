//ff:func feature=cli type=test control=sequence
//ff:what collectIssues test — Report ERROR/WARNING 수집 검증

package main

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func TestCollectIssues_E(t *testing.T) {
	t.Run("mpty", func(t *testing.T) {
		r := &validate.Report{}
		issues := collectIssues(r)
		if len(issues) != 0 {
			t.Fatalf("expected 0 issues, got %d", len(issues))
		}
	})
	t.Run("rrorsBeforeWarnings", func(t *testing.T) {
		r := &validate.Report{
			Steps: []validate.StepResult{
				{
					Diagnostics: []diagnostic.Diagnostic{
						{Level: diagnostic.LevelWarning, Message: "w1"},
						{Level: diagnostic.LevelError, Message: "e1"},
					},
				},
				{
					Diagnostics: []diagnostic.Diagnostic{
						{Level: diagnostic.LevelError, Message: "e2"},
					},
				},
			},
		}
		issues := collectIssues(r)
		if len(issues) != 3 {
			t.Fatalf("expected 3 issues, got %d", len(issues))
		}
		// errors should come before warnings
		if issues[0].Level != diagnostic.LevelError {
			t.Errorf("expected first issue to be ERROR, got %s", issues[0].Level)
		}
		if issues[2].Level != diagnostic.LevelWarning {
			t.Errorf("expected last issue to be WARNING, got %s", issues[2].Level)
		}
	})
}
