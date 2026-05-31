//ff:func feature=validate type=test control=iteration dimension=1
//ff:what TestXos67ResponseFieldType_Match — 서브테스트 디스패치
package openapi_ssac

import "testing"

func TestXos67ResponseFieldType_Match(t *testing.T) {
	for _, st := range []struct {
		name string
		fn   func(*testing.T)
	}{
		{"type match passes", subtestTestXos67ResponseFieldTypeMatchTypeMatchPasses},
		{"type mismatch raises diagnostic", subtestTestXos67ResponseFieldTypeMatchTypeMismatchRaisesDiagnostic},
		{"unresolvable expected type skipped", subtestTestXos67ResponseFieldTypeMatchUnresolvableExpectedTypeSkipped},
		{"timestamptz bound to date-time field passes", subtestTestXos67ResponseFieldTypeMatchTimestamptzBoundToDateTimeFieldPasses},
		{"string literal bound to date-time field errors (false negative closed)", subtestTestXos67ResponseFieldTypeMatchStringLiteralBoundToDateTimeFieldErrorsFalseNegativeClosed},
		{"string literal bound to string field passes", subtestTestXos67ResponseFieldTypeMatchStringLiteralBoundToStringFieldPasses},
		{"unresolvable actual type skipped", subtestTestXos67ResponseFieldTypeMatchUnresolvableActualTypeSkipped},
		{"uuid field bound to DB UUID column passes", subtestTestXos67ResponseFieldTypeMatchUuidFieldBoundToDBUUIDColumnPasses},
		{"uuid field bound to func openapi_types.UUID passes", subtestTestXos67ResponseFieldTypeMatchUuidFieldBoundToFuncOpenapiTypesUUIDPasses},
		{"uuid field bound to func pgtype.UUID errors with expected openapi_types.UUID", subtestTestXos67ResponseFieldTypeMatchUuidFieldBoundToFuncPgtypeUUIDErrorsWithExpectedOpenapiTypesUUID},
		{"uuid field bound to string literal errors (false negative closed)", subtestTestXos67ResponseFieldTypeMatchUuidFieldBoundToStringLiteralErrorsFalseNegativeClosed},
	} {
		t.Run(st.name, st.fn)
	}
}
