//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what findArgByKey -- 테스트용 FieldArg Key 검색 헬퍼

package ir

// findArgByKey returns the FieldArg with the given key, or nil.
func findArgByKey(args []FieldArg, key string) *FieldArg {
	for i := range args {
		if args[i].Key == key {
			return &args[i]
		}
	}
	return nil
}
