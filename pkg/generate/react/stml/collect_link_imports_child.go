//ff:func feature=stml-gen type=util control=selection topic=import-collect
//ff:what 단일 ChildNode를 Kind별로 분기하여 링크 임포트 플래그를 설정한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectLinkImportsChild handles one ChildNode for collectLinkImports:
// link nodes (and data-each RowLinks) flag the Link import, route.*
// sources flag useParams, container kinds recurse.
func collectLinkImportsChild(ch stmlparser.ChildNode, is *importSet) {
	switch ch.Kind {
	case "link":
		markLinkImports(*ch.Link, is)
		collectLinkImports(ch.Link.Children, is)
	case "fetch":
		collectLinkImports(ch.Fetch.Children, is)
	case "each":
		if ch.Each.RowLink != nil {
			markLinkImports(*ch.Each.RowLink, is)
		}
		collectLinkImports(ch.Each.Children, is)
	case "static":
		collectLinkImports(ch.Static.Children, is)
	case "state":
		collectLinkImports(ch.State.Children, is)
	}
}
