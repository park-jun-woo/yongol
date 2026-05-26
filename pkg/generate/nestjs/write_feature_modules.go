//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what writeFeatureModules — feature별 service+controller+module 파일 일괄 기록

package nestjs

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeFeatureModules renders and writes service + controller + module
// for each feature. Returns the list of feature names written.
func writeFeatureModules(plansByFeature map[string][]*ir.ServicePlan, srcDir string, reg ir.TypeRegistry) ([]string, error) {
	featureNames := make([]string, 0, len(plansByFeature))
	for feature, plans := range plansByFeature {
		if err := writeOneFeature(feature, plans, srcDir, reg); err != nil {
			return nil, err
		}
		featureNames = append(featureNames, feature)
	}
	return featureNames, nil
}
