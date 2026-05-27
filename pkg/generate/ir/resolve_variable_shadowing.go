//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what resolveVariableShadowing -- Op 목록에서 동일 VarName 중복 선언 감지 → _result 접미사로 해소 + 후속 Op FieldArg.Source 갱신

package ir

import "strings"

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

		switch ops[i].Kind {
		case OpGet:
			if ops[i].Get != nil && ops[i].Get.VarName != "" {
				orig := ops[i].Get.VarName
				ops[i].Get.VarName = resolveVar(orig, declared)
				if ops[i].Get.VarName != orig {
					renames[orig] = ops[i].Get.VarName
				}
			}
		case OpPost:
			if ops[i].Post != nil && ops[i].Post.VarName != "" {
				orig := ops[i].Post.VarName
				ops[i].Post.VarName = resolveVar(orig, declared)
				if ops[i].Post.VarName != orig {
					renames[orig] = ops[i].Post.VarName
				}
			}
		case OpCall:
			if ops[i].Call != nil && ops[i].Call.ResultVar != "" {
				orig := ops[i].Call.ResultVar
				ops[i].Call.ResultVar = resolveVar(orig, declared)
				if ops[i].Call.ResultVar != orig {
					renames[orig] = ops[i].Call.ResultVar
				}
			}
		case OpVerifyPassword:
			if ops[i].VerifyPW != nil && ops[i].VerifyPW.ResultVar != "" {
				orig := ops[i].VerifyPW.ResultVar
				ops[i].VerifyPW.ResultVar = resolveVar(orig, declared)
				if ops[i].VerifyPW.ResultVar != orig {
					renames[orig] = ops[i].VerifyPW.ResultVar
				}
			}
		}
	}
}

// applyRenames updates FieldArg.Source and ResponseField.Source references
// in an Op to match renamed variables from previous shadowing resolution.
func applyRenames(op *Op, renames map[string]string) {
	if len(renames) == 0 {
		return
	}
	for _, args := range collectFieldArgSlices(op) {
		for j := range *args {
			fa := &(*args)[j]
			if newName, ok := renames[fa.Source]; ok {
				fa.Source = newName
			}
		}
	}
	// ResponseOp fields also reference variables via dot-notation.
	if op.Kind == OpResponse && op.Response != nil {
		for j := range op.Response.Fields {
			f := &op.Response.Fields[j]
			dotIdx := strings.IndexByte(f.Source, '.')
			if dotIdx >= 0 {
				varPart := f.Source[:dotIdx]
				if newName, ok := renames[varPart]; ok {
					f.Source = newName + f.Source[dotIdx:]
				}
			} else {
				if newName, ok := renames[f.Source]; ok {
					f.Source = newName
				}
			}
		}
	}
}

// resolveVar checks whether varName collides with an already-declared
// variable. If so, it appends "_result" and retries until unique.
// The final name is registered in declared.
func resolveVar(varName string, declared map[string]bool) string {
	name := varName
	for declared[name] {
		name = name + "_result"
	}
	declared[name] = true
	return name
}
