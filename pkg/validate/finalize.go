//ff:func feature=validate type=util control=iteration dimension=1
//ff:what finalize — diagnostics에 ERROR 존재 여부로 StepResult 상태 결정
package validate

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

func finalize(name string, diags []diagnostic.Diagnostic) StepResult {
	status := StatusPass
	for _, d := range diags {
		if d.Level == diagnostic.LevelError {
			status = StatusFail
			break
		}
	}
	return StepResult{Name: name, Status: status, Diagnostics: diags}
}
