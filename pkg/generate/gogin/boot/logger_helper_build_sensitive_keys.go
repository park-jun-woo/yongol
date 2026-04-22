//ff:func feature=gen-gogin type=generator control=sequence
//ff:what loggerHelperBuildSensitiveKeys — redact.DefaultKeys + @sensitive 컬럼을 병합한 map 생성 헬퍼 소스 반환

package boot

// loggerHelperBuildSensitiveKeys returns the top-level buildSensitiveKeys(extras []string)
// helper source. Extracted from main() so that main no longer carries a depth-1 range loop
// (filefunc A13: selection + loop mixed at depth 1 is forbidden). The helper merges
// redact.DefaultKeys with DDL @sensitive columns supplied via extras.
func loggerHelperBuildSensitiveKeys() string {
	return `func buildSensitiveKeys(extras []string) map[string]bool {
	out := make(map[string]bool, len(redact.DefaultKeys)+len(extras))
	for k, v := range redact.DefaultKeys {
		out[k] = v
	}
	for _, k := range extras {
		out[k] = true
	}
	return out
}`
}
