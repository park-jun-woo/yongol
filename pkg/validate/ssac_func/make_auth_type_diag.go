//ff:func feature=validate type=rule control=sequence topic=func-check
//ff:what makeAuthTypeDiag — @auth input 의 sourceType 이 string 비호환이면 XFS-70 진단 생성

package ssac_func

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// makeAuthTypeDiag returns a diagnostic if sourceType is non-empty and not
// string-compatible, nil otherwise.
func makeAuthTypeDiag(fileName string, line int, inputKey, sourceType string) *diagnostic.Diagnostic {
	if sourceType == "" || TypesCompatible(sourceType, "string") {
		return nil
	}
	return &diagnostic.Diagnostic{
		File:  fileName,
		Line:  line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: "[XFS-70] @auth input " + inputKey +
			" value type " + sourceType + " is not string-compatible" +
			" (authz.CheckRequest." + inputKey + " is string)",
		Advice: "Use a string-typed source (e.g. request.id for path params) " +
			"instead of a DB row UUID field",
	}
}
