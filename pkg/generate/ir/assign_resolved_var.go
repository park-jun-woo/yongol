//ff:func feature=gen-ir type=util control=sequence
//ff:what assignResolvedVar -- 변수명 포인터를 충돌 없는 이름으로 갱신하고 변경 시 rename 등록

package ir

// assignResolvedVar resolves *varName against the declared set and, when the
// resolution changed the name, records the mapping in renames.
func assignResolvedVar(varName *string, declared map[string]bool, renames map[string]string) {
	if *varName == "" {
		return
	}
	orig := *varName
	*varName = resolveVar(orig, declared)
	if *varName != orig {
		renames[orig] = *varName
	}
}
