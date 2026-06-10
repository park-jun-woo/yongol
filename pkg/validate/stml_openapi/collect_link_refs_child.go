//ff:func feature=validate type=util control=selection topic=stml-openapi
//ff:what 단일 ChildNode를 Kind별로 분기하여 링크 참조를 수집한다 (each 진입 시 item 스키마 해석)

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectLinkRefsChild handles one ChildNode for collectLinkRefs. A
// data-fetch resets the context (no row in scope); a data-each resolves
// its item schema for the row context and records the RowLink, if any.
func collectLinkRefsChild(ch stml.ChildNode, opID string, itemFields map[string]bool, inEach bool, raif map[string]map[string]map[string]bool, out *[]linkRefCtx) {
	switch ch.Kind {
	case "link":
		*out = append(*out, linkRefCtx{Link: ch.Link, ItemFields: itemFields, InEach: inEach})
		collectLinkRefs(ch.Link.Children, opID, itemFields, inEach, raif, out)
	case "fetch":
		collectLinkRefs(ch.Fetch.Children, ch.Fetch.OperationID, nil, false, raif, out)
	case "each":
		var fields map[string]bool
		if byField, ok := raif[opID]; ok {
			fields = byField[ch.Each.Field]
		}
		if ch.Each.RowLink != nil {
			*out = append(*out, linkRefCtx{Link: ch.Each.RowLink, ItemFields: fields, InEach: true})
		}
		collectLinkRefs(ch.Each.Children, opID, fields, true, raif, out)
	case "static":
		if ch.Static != nil {
			collectLinkRefs(ch.Static.Children, opID, itemFields, inEach, raif, out)
		}
	case "state":
		if ch.State != nil {
			collectLinkRefs(ch.State.Children, opID, itemFields, inEach, raif, out)
		}
	}
}
