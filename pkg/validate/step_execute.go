//ff:func feature=validate type=util control=sequence
//ff:what step.execute — missing/wired 순으로 정책 판단 후 StepResult 반환
package validate

import "github.com/park-jun-woo/yongol/pkg/yongol"

// execute applies missing/run policy and returns a StepResult.
func (s step) execute(fs *yongol.Fullstack, cfg *config) StepResult {
	if !s.isPresent(fs) {
		return StepResult{Name: s.Name, Status: StatusMissing, Summary: "SSOT not detected"}
	}
	if s.Run == nil {
		return StepResult{Name: s.Name, Status: StatusSkip, Summary: "validator not wired"}
	}
	diags := s.run(fs)
	return finalize(s.Name, diags)
}
