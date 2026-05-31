//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what resolveVar -- 변수명이 declared 와 충돌하면 _result 접미사로 유일화 후 등록

package ir

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
