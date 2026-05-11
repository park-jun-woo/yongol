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
	// Top-level actions → nil (not inside any fetch)
	for _, a := range page.Actions {
		if _, ok := m[a.OperationID]; !ok {
			m[a.OperationID] = nil
		}
	}
	// Walk top-level children
	walkChildrenForFetchMap(page.Children, nil, m)
	return m
}

// walkChildrenForFetchMap recursively walks ChildNode trees. parentFetchOps
// contains the OperationIDs of ancestor fetch blocks.
func walkChildrenForFetchMap(nodes []stmlparser.ChildNode, parentFetchOps []string, m map[string][]string) {
	for _, ch := range nodes {
		switch ch.Kind {
		case "action":
			if _, ok := m[ch.Action.OperationID]; !ok {
				if len(parentFetchOps) > 0 {
					m[ch.Action.OperationID] = append([]string{}, parentFetchOps...)
				}
				// else: top-level, leave nil
			}
		case "fetch":
			// Enter fetch context: actions inside inherit this fetch's ops
			fetchOps := collectFetchOps(*ch.Fetch, nil)
			inner := append(append([]string{}, parentFetchOps...), fetchOps...)
			walkChildrenForFetchMap(ch.Fetch.Children, inner, m)
		case "state":
			walkChildrenForFetchMap(ch.State.Children, parentFetchOps, m)
		case "static":
			walkChildrenForFetchMap(ch.Static.Children, parentFetchOps, m)
		case "each":
			walkChildrenForFetchMap(ch.Each.Children, parentFetchOps, m)
		}
	}
}
