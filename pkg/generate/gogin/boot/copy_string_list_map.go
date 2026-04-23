//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what copyStringListMap — map[string][]string 의 deep copy (generator shared-state 보호)

package boot

// copyStringListMap deep-copies a map[string][]string so the generator
// never mutates the shared manifest instance.
func copyStringListMap(in map[string][]string) map[string][]string {
	if in == nil {
		return nil
	}
	out := make(map[string][]string, len(in))
	for k, v := range in {
		copied := make([]string, len(v))
		copy(copied, v)
		out[k] = copied
	}
	return out
}
