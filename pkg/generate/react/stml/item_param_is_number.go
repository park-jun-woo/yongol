//ff:func feature=stml-gen type=util control=sequence
//ff:what item.* 파라미터 소스가 응답 스키마상 이미 숫자 타입인지 확인한다
package stml

import (
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// itemParamIsNumber reports whether the param source is an item.<Field>
// reference whose field is already numeric ("integer" or "number") in the
// enclosing data-each item schema. Such values need no Number(...) wrapping
// even when the bound OpenAPI path parameter is an integer.
func itemParamIsNumber(p stmlparser.ParamBind, itemFieldTypes map[string]string) bool {
	if !strings.HasPrefix(p.Source, "item.") {
		return false
	}
	field := strings.TrimPrefix(p.Source, "item.")
	if i := strings.IndexByte(field, '.'); i >= 0 {
		field = field[:i]
	}
	typ := itemFieldTypes[field]
	return typ == "integer" || typ == "number"
}
