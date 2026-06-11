//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what tm53EachBinds — data-each 내부 data-bind을 항목 스키마 타입과 대조해 TM-53 검사

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm53EachBinds applies the TM-53 checks to each data-bind inside a data-each
// block, resolving the field type from the array item schema (itemFields:
// item field name → OpenAPI type). Item fields absent from the schema stay
// silent — TM-07/TM-30 own those.
func tm53EachBinds(e stml.EachBlock, opID, file string, itemFields map[string]string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, b := range e.Binds {
		typ, ok := itemFields[b.Name]
		if !ok {
			continue
		}
		diags = append(diags, tm53CheckBind(b, typ, opID, file)...)
	}
	return diags
}
