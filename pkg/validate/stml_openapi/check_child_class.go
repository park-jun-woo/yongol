//ff:func feature=validate type=rule control=selection topic=stml-openapi
//ff:what checkChildClass — ChildNode 타입별 class 검사 디스패치

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// checkChildClass dispatches class checks on a ChildNode.
func checkChildClass(cn stml.ChildNode, file string) []diagnostic.Diagnostic {
	switch cn.Kind {
	case "static":
		if cn.Static != nil {
			return checkStaticClass(*cn.Static, file)
		}
	case "fetch":
		if cn.Fetch != nil {
			return checkFetchClass(*cn.Fetch, file)
		}
	case "action":
		if cn.Action != nil {
			return checkActionClass(*cn.Action, file)
		}
	case "each":
		if cn.Each != nil {
			return checkEachClass(*cn.Each, file)
		}
	case "state":
		if cn.State != nil {
			return checkStateClass(*cn.State, file)
		}
	case "component":
		if cn.Component != nil {
			return emitIfClass(file, "ComponentRef", cn.Component.Name, cn.Component.ClassName)
		}
	case "bind":
		if cn.Bind != nil {
			return emitIfClass(file, "FieldBind", cn.Bind.Name, cn.Bind.ClassName)
		}
	}
	return nil
}
