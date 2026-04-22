//ff:func feature=orchestrator type=util control=iteration dimension=1
//ff:what DetectedSSOT 슬라이스를 SSOTKind 맵으로 변환하고 Presence 반영
package yongol

// indexDetected builds a kind→DetectedSSOT map while populating fs.Presences.
// Separated from ParseAll to keep the orchestrator as a flat sequence (Q1).
func indexDetected(detected []DetectedSSOT, fs *Fullstack) map[SSOTKind]DetectedSSOT {
	has := make(map[SSOTKind]DetectedSSOT)
	for _, d := range detected {
		has[d.Kind] = d
		fs.Presences[d.Kind] = d.Presence
	}
	return has
}
