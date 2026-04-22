//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-openapi
//ff:what toSet — []string → map[string]bool 변환 헬퍼

package openapi_ssac

// toSet converts a slice of strings into a set for O(1) lookups.
func toSet(xs []string) map[string]bool {
	out := make(map[string]bool, len(xs))
	for _, x := range xs {
		out[x] = true
	}
	return out
}
