//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what pickConvertRHS — convert<X> struct literal 의 오른쪽 표현식 선택 (JSONB/pgtype/cast/required 분기)

package ssac

// pickConvertRHS chooses the right-hand side for one struct-literal line.
// JSONB fields read from the pre-unmarshalled local map variable;
// api-typed fields (enum or format-string wrappers like
// openapi_types.Email) cast through the api-side named type since sqlc
// keeps these columns as plain string; required scalars assign the sqlc
// row value directly; optional scalars wrap with ptrOf so the *T api
// slot accepts them.
//
// Phase005 (pgx/v5 refit) — when sqlc pgx/v5 emits a pgtype wrapper for
// the column (pgtype.Timestamp, pgtype.UUID, pgtype.Text, pgtype.Numeric)
// the rhs is unwrapped first via pgtypeRowUnwrap before any apiCast /
// ptrOf wrapping. Primitive columns continue to take the direct path.
func pickConvertRHS(jsonName, apiField, dbField string, isRequired bool, jsonbs []jsonbFieldAlias, apiCast string, rowKind pgtypeRowKind) string {
	for _, j := range jsonbs {
		if j.jsonName == jsonName {
			return j.localVar
		}
	}
	base := "row." + dbField
	if unwrapped, ok := pgtypeRowUnwrap(rowKind, base, apiCast); ok {
		base = unwrapped
	}
	if apiCast != "" {
		if isRequired {
			return apiCast + "(" + base + ")"
		}
		return "ptrOf(" + apiCast + "(" + base + "))"
	}
	if isRequired {
		return base
	}
	return "ptrOf(" + base + ")"
}
