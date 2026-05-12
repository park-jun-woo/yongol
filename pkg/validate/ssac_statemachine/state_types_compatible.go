//ff:func feature=validate type=util control=sequence topic=ssac-statemachine
//ff:what stateTypesCompatible — actual Go 타입이 expected 에 대입 가능한지 판별

package ssac_statemachine

import "strings"

// stateTypesCompatible reports whether actual can be assigned to expected.
// Mirrors ssac_func.TypesCompatible logic.
func stateTypesCompatible(actual, expected string) bool {
	a := strings.TrimPrefix(actual, "*")
	e := strings.TrimPrefix(expected, "*")
	if a == e {
		return true
	}
	return a == "nil" || e == "nil"
}
