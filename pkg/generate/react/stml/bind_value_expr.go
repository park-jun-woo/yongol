//ff:func feature=stml-gen type=generator control=selection
//ff:what bindValueExpr — 스키마 타입에 따라 data-bind 값의 JSX 표현식({...})을 생성한다

package stml

import (
	"fmt"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// bindValueExpr returns the JSX children expression for a bind value at the
// given access path (e.g. "buildingData.can_delete" or "item.url"), branching
// on the OpenAPI schema type/format:
//
//	boolean              → {path ? 'Yes' : 'No'}
//	string + date        → {path ? new Date(path).toLocaleDateString() : ''}
//	string + date-time   → {path ? new Date(path).toLocaleString() : ''}
//	integer | number     → {typeof path === 'number' ? path.toLocaleString() : path}
//	string / unknown     → {path}                       (fallback — byte-identical)
//
// The <img> case is NOT handled here: an image is a whole-element rewrite
// (src attribute, not children), so renderBindJSX / renderEachJSX emit it
// directly. With a zero FieldTypeInfo the fallback {path} reproduces the
// pre-Phase037 emission exactly (plans/gen/frontend Phase037, BUG-126).
func bindValueExpr(path string, info oapiparser.FieldTypeInfo) string {
	switch {
	case info.Type == "boolean":
		return fmt.Sprintf("{%s ? 'Yes' : 'No'}", path)
	case info.Type == "string" && info.Format == "date":
		return fmt.Sprintf("{%s ? new Date(%s).toLocaleDateString() : ''}", path, path)
	case info.Type == "string" && info.Format == "date-time":
		return fmt.Sprintf("{%s ? new Date(%s).toLocaleString() : ''}", path, path)
	case info.Type == "integer" || info.Type == "number":
		return fmt.Sprintf("{typeof %s === 'number' ? %s.toLocaleString() : %s}", path, path, path)
	default:
		return fmt.Sprintf("{%s}", path)
	}
}
