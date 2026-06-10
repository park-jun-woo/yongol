//ff:func feature=stml-gen type=util control=iteration dimension=1 topic=import-collect
//ff:what ChildNode 트리의 링크 노드에서 필요한 임포트(Link, useParams)를 수집한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectLinkImports walks the ChildNode tree and flags the imports the
// link nodes need: react-router-dom's Link component, and useParams when
// a link binds a route.<Name> source (page-flow Phase007).
func collectLinkImports(children []stmlparser.ChildNode, is *importSet) {
	for _, ch := range children {
		collectLinkImportsChild(ch, is)
	}
}
