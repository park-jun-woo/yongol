//ff:type feature=features type=model
//ff:what DiffResult — features.yaml diff 결과 (Added/Removed)

package features

import featparser "github.com/park-jun-woo/yongol/pkg/parser/features"

// DiffResult holds the result of comparing two feature sets.
type DiffResult struct {
	Added   []featparser.Feature // ops in newFeats but not in oldFeats
	Removed []featparser.Feature // ops in oldFeats but not in newFeats
}
