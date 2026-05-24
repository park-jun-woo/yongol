//ff:func feature=validate type=test control=sequence
//ff:what Report.HasFailure — StatusFail 존재 여부 검증 (empty/pass/fail/mixed)

package validate

import "testing"

func TestReportHasFailure(t *testing.T) {
	t.Run("empty report returns false", func(t *testing.T) {
		r := &Report{}
		if r.HasFailure() {
			t.Error("expected false")
		}
	})

	t.Run("all pass returns false", func(t *testing.T) {
		r := &Report{Steps: []StepResult{
			{Name: "a", Status: StatusPass},
			{Name: "b", Status: StatusSkip},
		}}
		if r.HasFailure() {
			t.Error("expected false")
		}
	})

	t.Run("has failure returns true", func(t *testing.T) {
		r := &Report{Steps: []StepResult{
			{Name: "a", Status: StatusPass},
			{Name: "b", Status: StatusFail},
		}}
		if !r.HasFailure() {
			t.Error("expected true")
		}
	})

	t.Run("single failure returns true", func(t *testing.T) {
		r := &Report{Steps: []StepResult{
			{Name: "x", Status: StatusFail},
		}}
		if !r.HasFailure() {
			t.Error("expected true")
		}
	})
}
