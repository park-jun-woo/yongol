//ff:func feature=gen-filefunc type=util control=sequence
//ff:what defaultSSOTEntries — optional.ssot 고정 키 맵 반환 (openapi/ddl/ssac 등)
package filefunc

// defaultSSOTEntries returns the fixed ssot keys allowed in annotations.
// These mirror the SSOT kinds supported by yongol.
func defaultSSOTEntries() map[string]string {
	return map[string]string{
		"openapi":  "",
		"ddl":      "",
		"ssac":     "",
		"states":   "",
		"policy":   "",
		"scenario": "",
		"funcspec": "",
		"config":   "",
	}
}
