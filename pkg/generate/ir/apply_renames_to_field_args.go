//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what applyRenamesToFieldArgs -- FieldArg 슬라이스의 Source 를 rename 맵으로 치환

package ir

// applyRenamesToFieldArgs rewrites each FieldArg.Source in args using renames.
func applyRenamesToFieldArgs(args []FieldArg, renames map[string]string) {
	for j := range args {
		if newName, ok := renames[args[j].Source]; ok {
			args[j].Source = newName
		}
	}
}
