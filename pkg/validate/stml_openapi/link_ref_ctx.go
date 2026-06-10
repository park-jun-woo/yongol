//ff:type feature=validate type=model topic=stml-openapi
//ff:what linkRefCtx — 수집된 data-link 참조와 둘러싼 each 컨텍스트(item 스키마)

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// linkRefCtx pairs a collected data-link reference with its enclosing
// data-each context, so TM-32 can resolve item.<Field> sources the same
// way TM-30 resolves item.* param sources.
type linkRefCtx struct {
	Link       *stml.LinkRef
	ItemFields map[string]bool // item schema of the innermost enclosing data-each (nil = unresolved)
	InEach     bool            // whether the link sits inside a data-each block
}
