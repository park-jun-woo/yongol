//ff:func feature=gen-filefunc type=util control=sequence
//ff:what defaultPatternEntries — optional.pattern 고정 키 맵 반환 (early-return 등)
package filefunc

// defaultPatternEntries returns the fixed pattern keys commonly used in
// generated backend code.
func defaultPatternEntries() map[string]string {
	return map[string]string{
		"early-return":     "",
		"error-collection": "",
	}
}
