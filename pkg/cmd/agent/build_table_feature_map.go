//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what buildTableFeatureMap — features를 table 이름으로 그룹화

package agent

import "github.com/park-jun-woo/yongol/pkg/parser/features"

// buildTableFeatureMap groups features by their table field.
func buildTableFeatureMap(ff *features.FeaturesFile) map[string][]features.Feature {
	m := make(map[string][]features.Feature)
	for _, f := range ff.Features {
		if f.Table != "" {
			m[f.Table] = append(m[f.Table], f)
		}
	}
	return m
}
