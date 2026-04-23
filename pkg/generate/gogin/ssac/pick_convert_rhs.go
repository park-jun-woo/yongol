//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what pickConvertRHS — convert<X> struct literal 의 오른쪽 표현식 선택 (JSONB/cast/required 분기)

package ssac

// pickConvertRHS chooses the right-hand side for one struct-literal line.
// JSONB fields read from the pre-unmarshalled local map variable;
// api-typed fields (enum or format-string wrappers like
// openapi_types.Email) cast through the api-side named type since sqlc
// keeps these columns as plain string; required scalars assign the sqlc
// row value directly; optional scalars wrap with ptrOf so the *T api
// slot accepts them.
func pickConvertRHS(jsonName, apiField, dbField string, isRequired bool, jsonbs []jsonbFieldAlias, apiCast string) string {
	for _, j := range jsonbs {
		if j.jsonName == jsonName {
			return j.localVar
		}
	}
	if apiCast != "" {
		if isRequired {
			return apiCast + "(row." + dbField + ")"
		}
		return "ptrOf(" + apiCast + "(row." + dbField + "))"
	}
	if isRequired {
		return "row." + dbField
	}
	return "ptrOf(row." + dbField + ")"
}
