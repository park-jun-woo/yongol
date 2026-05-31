//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what nonFeatureExternalModules — external package 중 feature module 이 아닌 것들의 이름 목록

package nestjs

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// nonFeatureExternalModules returns the names of external packages that are not
// also feature modules, so they can be registered as infra modules without
// duplication.
func nonFeatureExternalModules(extPkgs []externalPackage, plansByFeature map[string][]*ir.ServicePlan) []string {
	var names []string
	for _, ep := range extPkgs {
		if _, isFeature := plansByFeature[ep.Name]; !isFeature {
			names = append(names, ep.Name)
		}
	}
	return names
}
