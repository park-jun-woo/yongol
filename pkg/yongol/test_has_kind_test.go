//ff:func feature=orchestrator type=test-helper control=iteration dimension=1
//ff:what hasKind — DetectedSSOT 슬라이스에서 특정 Kind 를 탐색
package yongol

// hasKind reports whether detected contains a DetectedSSOT for k.
func hasKind(detected []DetectedSSOT, k SSOTKind) (DetectedSSOT, bool) {
	for _, d := range detected {
		if d.Kind == k {
			return d, true
		}
	}
	return DetectedSSOT{}, false
}
