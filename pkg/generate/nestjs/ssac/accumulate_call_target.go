//ff:func feature=gen-nestjs type=util control=sequence
//ff:what accumulateCallTarget — @call op 의 TargetFeature 를 cross-feature 또는 same-feature-stub 로 분류

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// accumulateCallTarget classifies an op's @call target as a cross-feature
// dependency or a same-feature stub requirement.
func accumulateCallTarget(op ir.Op, lowerFeature string, deps *moduleDeps, callFeatures map[string]bool) {
	if op.Kind != ir.OpCall || op.Call == nil || op.Call.TargetFeature == "" {
		return
	}
	if op.Call.TargetFeature != lowerFeature {
		callFeatures[op.Call.TargetFeature] = true
		return
	}
	deps.NeedsSameFeatureStub = true
}
