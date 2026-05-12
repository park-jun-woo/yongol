//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what ChildNode 트리를 재귀 탐색하여 action-fetch 매핑을 수집한다

package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// walkChildrenForFetchMap recursively walks ChildNode trees. parentFetchOps
// contains the OperationIDs of ancestor fetch blocks.
func walkChildrenForFetchMap(nodes []stmlparser.ChildNode, parentFetchOps []string, m map[string][]string) {
	for _, ch := range nodes {
		switch ch.Kind {
		case "action":
			recordActionFetchMapping(ch.Action.OperationID, parentFetchOps, m)
		case "fetch":
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
