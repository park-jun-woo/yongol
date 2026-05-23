//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what copyMap — map[string]any 얕은 복사

package agent

func copyMap(m map[string]any) map[string]any {
	out := make(map[string]any, len(m))
	for k, v := range m {
		out[k] = v
	}
	return out
}
