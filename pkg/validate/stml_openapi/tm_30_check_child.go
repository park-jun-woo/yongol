//ff:func feature=validate type=rule control=selection topic=stml-openapi
//ff:what TM-30 보조 — 단일 ChildNode를 each 컨텍스트(item 스키마)에 따라 분기 검사

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm30CheckChild dispatches one ChildNode carrying the enclosing context:
// opID is the nearest data-fetch operationId (route of the item schema
// lookup), itemFields is the item schema of the innermost enclosing
// data-each (nil when unresolved), inEach tells whether the node sits
// inside a data-each. A data-fetch resets the context: the parser never
// places a fetch inside a data-each, so its params are always row-less.
func tm30CheckChild(ch stml.ChildNode, file, opID string, itemFields map[string]bool, inEach bool, raif map[string]map[string]map[string]bool, out *[]diagnostic.Diagnostic) {
	switch ch.Kind {
	case "action":
		*out = append(*out, tm30CheckActionParams(*ch.Action, file, itemFields, inEach)...)
	case "fetch":
		*out = append(*out, tm30CheckFetchParams(*ch.Fetch, file)...)
		tm30CheckChildren(ch.Fetch.Children, file, ch.Fetch.OperationID, nil, false, raif, out)
	case "each":
		var fields map[string]bool
		if byField, ok := raif[opID]; ok {
			fields = byField[ch.Each.Field]
		}
		tm30CheckChildren(ch.Each.Children, file, opID, fields, true, raif, out)
	case "static":
		if ch.Static != nil {
			tm30CheckChildren(ch.Static.Children, file, opID, itemFields, inEach, raif, out)
		}
	case "state":
		if ch.State != nil {
			tm30CheckChildren(ch.State.Children, file, opID, itemFields, inEach, raif, out)
		}
	}
}
