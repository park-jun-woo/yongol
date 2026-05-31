//ff:func feature=validate type=test control=iteration dimension=1
//ff:what TestXdo77PropDiag — 서브테스트 디스패치
package openapi_ddl

import "testing"

func TestXdo77PropDiag(t *testing.T) {
	for _, st := range []struct {
		name string
		fn   func(*testing.T)
	}{
		{"nil propRef returns false", subtestTestXdo77PropDiagNilPropRefReturnsFalse},
		{"column not in DDL returns false", subtestTestXdo77PropDiagColumnNotInDDLReturnsFalse},
		{"matching types returns false", subtestTestXdo77PropDiagMatchingTypesReturnsFalse},
		{"type mismatch returns diagnostic", subtestTestXdo77PropDiagTypeMismatchReturnsDiagnostic},
		{"unknown DDL Go type returns false", subtestTestXdo77PropDiagUnknownDDLGoTypeReturnsFalse},
		{"empty type slice returns false", subtestTestXdo77PropDiagEmptyTypeSliceReturnsFalse},
		{"format mismatch with format in display", subtestTestXdo77PropDiagFormatMismatchWithFormatInDisplay},
		{"boolean match returns false", subtestTestXdo77PropDiagBooleanMatchReturnsFalse},
		{"float64 number with format double returns false", subtestTestXdo77PropDiagFloat64NumberWithFormatDoubleReturnsFalse},
		{"float64 formatless number returns float-specific diagnostic", subtestTestXdo77PropDiagFloat64FormatlessNumberReturnsFloatSpecificDiagnostic},
		{"float64 number with format float returns float-specific diagnostic", subtestTestXdo77PropDiagFloat64NumberWithFormatFloatReturnsFloatSpecificDiagnostic},
		{"float64 wrong type returns generic diagnostic", subtestTestXdo77PropDiagFloat64WrongTypeReturnsGenericDiagnostic},
	} {
		t.Run(st.name, st.fn)
	}
}
