//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-30 보조 — ChildNode 슬라이스를 동일 컨텍스트로 순회

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm30CheckChildren walks a ChildNode slice under one enclosing context.
func tm30CheckChildren(children []stml.ChildNode, file, opID string, itemFields map[string]bool, inEach bool, raif map[string]map[string]map[string]bool, out *[]diagnostic.Diagnostic) {
	for _, ch := range children {
		tm30CheckChild(ch, file, opID, itemFields, inEach, raif, out)
	}
}
