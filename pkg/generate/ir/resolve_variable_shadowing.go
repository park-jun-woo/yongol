//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what resolveVariableShadowing -- Op 목록에서 동일 VarName 중복 선언 감지 → _result 접미사로 해소

package ir

// resolveVariableShadowing detects duplicate variable declarations across
// ops and renames collisions with a _result suffix. Only renames when an
// earlier op already declared the same variable name — SSaC author's
// chosen names are preserved on first use. This resolves shadowing at
// build time so renderers need not handle it.
func resolveVariableShadowing(ops []Op) {
	declared := make(map[string]bool)
	for i := range ops {
		switch ops[i].Kind {
		case OpGet:
			if ops[i].Get != nil && ops[i].Get.VarName != "" {
				ops[i].Get.VarName = resolveVar(ops[i].Get.VarName, declared)
			}
		case OpPost:
			if ops[i].Post != nil && ops[i].Post.VarName != "" {
				ops[i].Post.VarName = resolveVar(ops[i].Post.VarName, declared)
			}
		case OpCall:
			if ops[i].Call != nil && ops[i].Call.ResultVar != "" {
				ops[i].Call.ResultVar = resolveVar(ops[i].Call.ResultVar, declared)
			}
		case OpVerifyPassword:
			if ops[i].VerifyPW != nil && ops[i].VerifyPW.ResultVar != "" {
				ops[i].VerifyPW.ResultVar = resolveVar(ops[i].VerifyPW.ResultVar, declared)
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
