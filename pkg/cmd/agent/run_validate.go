//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what runValidate — DetectSSOTs → ParseAll → Validate 실행 후 전체 진단 반환

package agent

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// runValidate runs DetectSSOTs → ParseAll → Validate and returns all diagnostics.
func runValidate(specsDir string) ([]diagnostic.Diagnostic, error) {
	detected, err := yongol.DetectSSOTs(specsDir)
	if err != nil {
		return nil, fmt.Errorf("detect SSOTs: %w", err)
	}
	fs := yongol.ParseAll(specsDir, detected)
	if len(fs.ParseDiagnostics) > 0 {
		return fs.ParseDiagnostics, nil
	}
	report := validate.Validate(fs)
	var all []diagnostic.Diagnostic
	for _, step := range report.Steps {
		all = append(all, step.Diagnostics...)
	}
	return all, nil
}
