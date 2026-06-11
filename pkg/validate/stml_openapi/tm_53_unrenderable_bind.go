//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-53 — data-bind가 렌더 불가/강등되는 경우 경고 (비스칼라·미지원 태그·img 타입 불일치)

package stml_openapi

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm53UnrenderableBind warns (WARNING) about data-bind fields that cannot
// render as readable content (plans/gen/frontend Phase037, BUG-126):
//
//	(a) non-scalar  — an object/array field bound as text → "[object Object]"
//	                  or a comma-joined string; use a dotted path / data-each.
//	(b) bad tag     — data-bind on a void/media tag codegen cannot bind
//	                  (everything except <img>).
//	(c) img mismatch— <img data-bind> whose field is not a string URL.
//
// boolean is intentionally NOT warned: the codegen now renders it as Yes/No.
// Direct fetch binds consult the response schema (responseFields); data-each
// binds consult the array item schema (itemTypes, via tm53EachBinds). Unknown
// fields stay silent — TM-06/07 own those. Dotted binds are skipped:
// responseFields exposes only top-level types, so the leaf type is unknown.
func tm53UnrenderableBind(f stml.FetchBlock, opID, file string, entry operationEntry, itemTypes map[string]map[string]map[string]string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	respFields := responseFields(entry.op)
	for _, b := range f.Binds {
		if strings.IndexByte(b.Name, '.') >= 0 {
			continue
		}
		info, ok := respFields[b.Name]
		if !ok {
			continue
		}
		diags = append(diags, tm53CheckBind(b, info.typ, opID, file)...)
	}
	for _, e := range f.Eaches {
		diags = append(diags, tm53EachBinds(e, opID, file, itemTypes[opID][e.Field])...)
	}
	return diags
}
