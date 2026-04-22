//ff:func feature=ssac-parse type=parser control=sequence
//ff:what 결과 변수가 있는 CRUD 시퀀스(Type var = Model.Method(...))를 파싱한다

package ssac

import "strings"

// parseCRUDWithResult parses a CRUD sequence that has an assignment result.
// Format: Type var = Model.Method({Key: val, ...})
func parseCRUDWithResult(rest string, seq *Sequence) error {
	eqIdx := strings.Index(rest, "=")
	if eqIdx < 0 {
		return nil
	}
	lhs := strings.TrimSpace(rest[:eqIdx])
	rhs := strings.TrimSpace(rest[eqIdx+1:])

	result := parseResult(lhs)
	if result == nil {
		return nil
	}
	seq.Result = result

	model, inputs, _, err := parseCallExprInputs(rhs)
	if err != nil {
		return err
	}
	seq.Package, seq.Model = splitPackagePrefix(model)
	seq.Inputs = inputs
	return nil
}
