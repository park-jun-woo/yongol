//ff:func feature=features type=util control=iteration dimension=1
//ff:what DiffOps — features.yaml 간 operationId 집합 diff (신규/삭제 추출)

package features

import (
	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
)

// DiffResult holds the result of comparing two feature sets.
type DiffResult struct {
	Added   []featparser.Feature // ops in newFeats but not in oldFeats
	Removed []featparser.Feature // ops in oldFeats but not in newFeats
}

// DiffOps compares two feature slices by operationId and returns added/removed.
func DiffOps(oldFeats, newFeats []featparser.Feature) DiffResult {
	oldSet := make(map[string]featparser.Feature, len(oldFeats))
	for _, f := range oldFeats {
		oldSet[f.Op] = f
	}

	newSet := make(map[string]featparser.Feature, len(newFeats))
	for _, f := range newFeats {
		newSet[f.Op] = f
	}

	var result DiffResult
	for _, f := range newFeats {
		if _, exists := oldSet[f.Op]; !exists {
			result.Added = append(result.Added, f)
		}
	}
	for _, f := range oldFeats {
		if _, exists := newSet[f.Op]; !exists {
			result.Removed = append(result.Removed, f)
		}
	}
	return result
}
