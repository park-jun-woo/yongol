//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what writeFeatureModules — feature별 service+router 파일 일괄 기록

package fastapi

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeFeatureModules renders and writes service + router for each feature.
// Returns the list of feature names written.
func writeFeatureModules(plansByFeature map[string][]*ir.ServicePlan, appDir string, reg ir.TypeRegistry) ([]string, error) {
	featureNames := make([]string, 0, len(plansByFeature))
	for feature, plans := range plansByFeature {
		if err := writeOneFeature(feature, plans, appDir, reg); err != nil {
			return nil, err
		}
		featureNames = append(featureNames, feature)
	}
	return featureNames, nil
}
