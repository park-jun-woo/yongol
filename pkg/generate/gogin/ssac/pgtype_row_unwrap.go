//ff:func feature=gen-gogin type=util control=selection
//ff:what pgtypeRowUnwrap — sqlc pgx/v5 row 필드의 언래핑 표현식 선택

package ssac

// pgtypeRowUnwrap returns the Go expression that extracts the primitive
// value from a sqlc pgx/v5 row field. rowAccess is the full selector (e.g.
// "row.CreatedAt"). Returns ("", false) when no unwrap is required (the
// primitive path — caller emits rowAccess verbatim).
func pgtypeRowUnwrap(kind pgtypeRowKind, rowAccess string, apiCast string) (string, bool) {
	switch kind {
	case pgTimestamp:
		return rowAccess + ".Time", true
	case pgTextWrapper:
		return rowAccess + ".String", true
	case pgUUID:
		// pgUUIDToString helper centralises the Valid + [16]byte → canonical
		// UUID string conversion. apiCast (e.g. openapi_types.UUID) is applied
		// by the caller on top.
		_ = apiCast
		return "pgUUIDToString(" + rowAccess + ")", true
	case pgNumeric:
		return "pgNumericToString(" + rowAccess + ")", true
	}
	return "", false
}
