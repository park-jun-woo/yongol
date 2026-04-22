//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what parseDottedValue — "source.field" 문자열을 (source, field) 로 분리

package ssac

// parseDottedValue splits "source.field" into (source, field). Returns
// ("", "") if v has no dot.
func parseDottedValue(v string) (string, string) {
	for i := 0; i < len(v); i++ {
		if v[i] == '.' {
			return v[:i], v[i+1:]
		}
	}
	return "", ""
}
