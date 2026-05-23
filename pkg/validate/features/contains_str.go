//ff:func feature=validate type=util control=iteration dimension=1 topic=features-structural
//ff:what containsStr — 문자열 슬라이스에 대상 문자열 포함 여부 확인

package features

// containsStr returns true if slice contains the target string.
func containsStr(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}
