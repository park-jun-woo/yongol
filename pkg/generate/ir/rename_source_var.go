//ff:func feature=gen-ir type=util control=sequence
//ff:what renameSourceVar -- dot-notation source 문자열의 변수 부분을 rename 맵으로 치환

package ir

import "strings"

// renameSourceVar returns source with its leading variable part rewritten via
// renames. When source contains a dot, only the part before the first dot is
// considered the variable; the accessor suffix is preserved.
func renameSourceVar(source string, renames map[string]string) string {
	dotIdx := strings.IndexByte(source, '.')
	if dotIdx < 0 {
		if newName, ok := renames[source]; ok {
			return newName
		}
		return source
	}
	varPart := source[:dotIdx]
	if newName, ok := renames[varPart]; ok {
		return newName + source[dotIdx:]
	}
	return source
}
