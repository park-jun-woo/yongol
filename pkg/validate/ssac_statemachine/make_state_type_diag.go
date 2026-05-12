//ff:func feature=validate type=rule control=sequence topic=ssac-statemachine
//ff:what makeStateTypeDiag — sourceType 이 string 비호환이면 XSM-71 진단 생성

package ssac_statemachine

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// makeStateTypeDiag returns a diagnostic if sourceType is non-empty and not
// string-compatible, nil otherwise.
func makeStateTypeDiag(fileName string, line int, inputKey, sourceType string) *diagnostic.Diagnostic {
	if sourceType == "" || stateTypesCompatible(sourceType, "string") {
		return nil
	}
	return &diagnostic.Diagnostic{
		File:  fileName,
		Line:  line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: "[XSM-71] @state input " + inputKey +
			" value type " + sourceType + " is not string-compatible" +
			" (statemachine parameter is string)",
		Advice: "Use a string-typed field (e.g. status TEXT column) " +
			"instead of a UUID or numeric column",
	}
}
