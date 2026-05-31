//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what resolveVariableShadowing -- Op 목록에서 동일 VarName 중복 선언 감지 → _result 접미사로 해소 + 후속 Op FieldArg.Source 갱신

package ir

// resolveVariableShadowing detects duplicate variable declarations across
// ops and renames collisions with a _result suffix. Only renames when an
// earlier op already declared the same variable name — SSaC author's
// chosen names are preserved on first use.
//
// After renaming a variable, all subsequent ops that reference the old
// variable name in FieldArg.Source or ResponseField.Source are updated
// to use the new name. This ensures rendered code accesses the correct
// variable (e.g. user_result.email instead of user.email).
//
// reservedNames seeds the declared set with names that are already taken by
// method parameters (e.g. "params", "body", "query", "user", "payload") so
// that op result variables do not shadow them.
func resolveVariableShadowing(ops []Op, reservedNames ...string) {
	declared := make(map[string]bool)
	for _, n := range reservedNames {
		declared[n] = true
	}
	renames := make(map[string]string)

	for i := range ops {
		// Apply renames from earlier ops to current op's FieldArg references.
		applyRenames(&ops[i], renames)
		renameOpResultVar(&ops[i], declared, renames)
	}
}
