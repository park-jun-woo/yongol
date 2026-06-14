//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what ChildNode 트리의 모든 ActionBlock Params에 optional route 플래그를 재귀로 표시한다 (BUG-136)
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// markChildActionBindsOptional walks the ChildNode tree and flags optional
// route params on every nested ActionBlock (including row actions inside
// data-each). It mirrors collectAllActions's traversal but mutates the binds
// in place through the ChildNode pointers.
func markChildActionBindsOptional(children []stmlparser.ChildNode, required map[string]bool) {
	for _, ch := range children {
		switch ch.Kind {
		case "action":
			setBindsOptional(ch.Action.Params, required)
		case "fetch":
			markChildActionBindsOptional(ch.Fetch.Children, required)
		case "state":
			markChildActionBindsOptional(ch.State.Children, required)
		case "static":
			markChildActionBindsOptional(ch.Static.Children, required)
		case "each":
			markChildActionBindsOptional(ch.Each.Children, required)
		}
	}
}
