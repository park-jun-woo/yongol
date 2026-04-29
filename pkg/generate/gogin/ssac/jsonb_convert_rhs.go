//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what jsonbConvertRHS — JSONB 로컬 변수 매칭 시 (rhs, true) 반환 (nullable 은 & 부착)

package ssac

// jsonbConvertRHS scans the pre-collected JSONB field aliases for a
// match on jsonName. When found:
//
//   - required column → return the local var value form (map[string]any)
//   - nullable column → return "&local" so the struct literal accepts
//     the *map[string]any field shape (BUG-038)
//
// Returns ("", false) when no JSONB alias matches; the caller then
// falls through to the type-binding path.
func jsonbConvertRHS(jsonName string, isRequired bool, jsonbs []jsonbFieldAlias) (string, bool) {
	for _, j := range jsonbs {
		if j.jsonName != jsonName {
			continue
		}
		if !isRequired {
			return "&" + j.localVar, true
		}
		return j.localVar, true
	}
	return "", false
}
