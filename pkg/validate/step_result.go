//ff:type feature=validate type=model
//ff:what StepResult — 단일 validator 실행 결과
package validate

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// StepResult captures one validator's result.
type StepResult struct {
	Name        string
	Status      Status
	Summary     string
	Diagnostics []diagnostic.Diagnostic
}
