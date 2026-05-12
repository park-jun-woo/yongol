//ff:func feature=generate type=util control=iteration dimension=1
//ff:what STML field-less operationId 집합을 NoBodyOps 맵에 병합한다

package generate

// mergeFieldlessOps copies all keys from src into dst.
func mergeFieldlessOps(dst, src map[string]bool) {
	for opID := range src {
		dst[opID] = true
	}
}
