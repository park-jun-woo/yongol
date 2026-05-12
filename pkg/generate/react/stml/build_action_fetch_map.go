//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what Action이 속한 FetchBlock의 OperationID를 매핑한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// buildActionFetchMap walks the page tree and maps each action's OperationID to
// the fetch OperationIDs it should invalidate on success.
//
// Rules:
//   - Action inside a data-fetch (directly or via data-state) → invalidate that
//     fetch's own OperationID only.
//   - Action at top level (not inside any fetch) → nil entry (caller falls back
//     to invalidating all page-level fetchOps).
func buildActionFetchMap(page stmlparser.PageSpec) map[string][]string {
	m := make(map[string][]string)
	for _, a := range page.Actions {
		if _, ok := m[a.OperationID]; !ok {
			m[a.OperationID] = nil
		}
	}
	walkChildrenForFetchMap(page.Children, nil, m)
	return m
}
