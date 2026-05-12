//ff:func feature=validate type=util control=selection topic=stml-design
//ff:what extractChildTokens — ChildNode를 종류별로 분기하여 토큰 추출
package stml_design

import (
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// extractChildTokens processes a ChildNode recursively.
func extractChildTokens(cn stml.ChildNode, file string, out *pageTokenRefs) {
	switch cn.Kind {
	case "static":
		if cn.Static != nil {
			extractStaticTokens(*cn.Static, file, out)
		}
	case "fetch":
		if cn.Fetch != nil {
			extractFetchTokens(*cn.Fetch, file, out)
		}
	case "action":
		if cn.Action != nil {
			extractActionTokens(*cn.Action, file, out)
		}
	case "each":
		if cn.Each != nil {
			extractEachTokens(*cn.Each, file, out)
		}
	case "component":
		if cn.Component != nil {
			recordComponent(*cn.Component, file, out)
			classifyTokens(cn.Component.ClassName, file, out)
		}
	case "bind":
		if cn.Bind != nil {
			classifyTokens(cn.Bind.ClassName, file, out)
		}
	}
}
