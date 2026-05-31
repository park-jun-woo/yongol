//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what applyRenamesToResponseFields -- ResponseField.Source(dot-notation 포함)를 rename 맵으로 치환

package ir

// applyRenamesToResponseFields rewrites each ResponseField.Source in fields
// using renames. Sources may carry dot-notation (e.g. "user.email"); only the
// variable part before the first dot is rewritten.
func applyRenamesToResponseFields(fields []ResponseField, renames map[string]string) {
	for j := range fields {
		fields[j].Source = renameSourceVar(fields[j].Source, renames)
	}
}
