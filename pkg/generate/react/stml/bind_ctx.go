//ff:type feature=stml-gen type=model
//ff:what bindCtx — 타입 인지 data-bind 렌더링용 응답 필드 타입 맵 + 현재 fetch op 스코프

package stml

import oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"

// bindCtx threads the response-field type map through the JSX renderers along
// with the current fetch scope (opID). The full map is carried — not a
// per-op slice — so a nested data-fetch re-scopes correctly: renderFetchJSX
// resets opID to its own f.OperationID before recursing, mirroring
// validateFetchBlock's per-op recursion. A zero bindCtx (all nil, opID "")
// makes field() return the zero FieldTypeInfo, so bindValueExpr falls back to
// the plain {value} emission — keeping option-unwired output byte-identical.
type bindCtx struct {
	all  map[string]map[string]oapiparser.FieldTypeInfo // operationId → field path → type/format
	opID string                                         // current fetch scope (re-set by renderFetchJSX)
}

// field returns the FieldTypeInfo for a bind field path within the current
// fetch scope, or the zero value when the map, op, or field is absent.
func (c bindCtx) field(name string) oapiparser.FieldTypeInfo {
	return c.all[c.opID][name]
}
