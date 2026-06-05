//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-openapi
//ff:what 페이지 fetch들의 top-level 응답 prop → fetch operationEntry 맵을 구성한다 (TM-14용)

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// buildModelFetchMap maps each guard model prefix (a top-level response property
// of a page fetch) to the operationEntry of the fetch that provides it. It
// mirrors how data-bind dotted paths resolve their top-level key against the
// fetch response schema (TM-06), reusing responseFields. Fetches whose
// operationId is unknown in opMap are skipped (TM-01 reports those).
func buildModelFetchMap(fetches []stml.FetchBlock, opMap map[string]operationEntry) map[string]operationEntry {
	out := make(map[string]operationEntry)
	for _, f := range fetches {
		entry, ok := opMap[f.OperationID]
		if !ok {
			continue
		}
		for prop := range responseFields(entry.op) {
			out[prop] = entry
		}
	}
	return out
}
