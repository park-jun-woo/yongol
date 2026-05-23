//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what buildFeatureLookupFromFF — FeaturesFile에서 op → Feature 맵 구성

package agent

import "github.com/park-jun-woo/yongol/pkg/parser/features"

func buildFeatureLookupFromFF(ff *features.FeaturesFile) map[string]features.Feature {
	if ff == nil {
		return nil
	}
	m := make(map[string]features.Feature, len(ff.Features))
	for _, f := range ff.Features {
		m[f.Op] = f
	}
	return m
}
