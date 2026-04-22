//ff:func feature=validate type=util control=selection topic=ddl-structural
//ff:what toStringSlice — YAML 값(string 또는 []string)을 []string 으로 정규화

package ddl

// toStringSlice normalises a YAML value that may be a single string or a list
// of strings into []string. Returns nil for unrecognised types. Loop extraction
// for the list branch lives in extractStringsFromSlice.
func toStringSlice(v interface{}) []string {
	switch val := v.(type) {
	case string:
		return []string{val}
	case []interface{}:
		return extractStringsFromSlice(val)
	}
	return nil
}
