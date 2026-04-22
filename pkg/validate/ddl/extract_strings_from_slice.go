//ff:func feature=validate type=util control=iteration dimension=1 topic=ddl-structural
//ff:what extractStringsFromSlice — []interface{} 에서 string 원소만 추출

package ddl

// extractStringsFromSlice filters items to strings only. Used by toStringSlice
// for the YAML list branch.
func extractStringsFromSlice(items []interface{}) []string {
	var out []string
	for _, item := range items {
		if s, ok := item.(string); ok {
			out = append(out, s)
		}
	}
	return out
}
