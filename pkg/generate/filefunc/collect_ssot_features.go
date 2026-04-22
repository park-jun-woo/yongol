//ff:func feature=gen-filefunc type=util control=sequence
//ff:what collectSSOTFeatures — SSaC / funcspec 에서 도메인 feature 후보 + 설명을 수집
package filefunc

import "github.com/park-jun-woo/yongol/pkg/yongol"

// collectSSOTFeatures scans Fullstack SSOT data and returns a map of
// feature-name → description. The feature name is derived from SSaC
// ServiceFunc.Feature (domain folder name) and funcspec.Package (support
// folder name). Descriptions come from funcspec.Description when available,
// otherwise left as empty string for the caller to fill in.
func collectSSOTFeatures(fs *yongol.Fullstack) map[string]string {
	out := map[string]string{}
	if fs == nil {
		return out
	}
	addServiceFuncFeatures(out, fs)
	addFuncSpecFeatures(out, fs)
	return out
}
