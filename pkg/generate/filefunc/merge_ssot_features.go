//ff:func feature=gen-filefunc type=util control=iteration dimension=1
//ff:what mergeSSOTFeatures — SSOT 에서 수집한 feature-설명 쌍을 타겟 맵에 덮어쓴다
package filefunc

// mergeSSOTFeatures copies SSOT-derived (name, desc) pairs into dst, using
// resolveFeatureDescription so that empty descriptions fall back to the
// infra baseline or fallback text.
func mergeSSOTFeatures(dst, ssot map[string]string) {
	for name, desc := range ssot {
		dst[name] = resolveFeatureDescription(name, desc)
	}
}
