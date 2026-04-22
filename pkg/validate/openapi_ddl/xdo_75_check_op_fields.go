//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-ddl
//ff:what xdo75CheckOpFields — 단일 operation 의 request field 를 검사하여 XDO-75 진단 생성

package openapi_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// xdo75CheckOpFields inspects the RequestConstraints map for a single
// operation and emits XDO-75 diagnostics for optional fields whose DDL
// column is NOT NULL without DEFAULT.
func xdo75CheckOpFields(fs *yongol.Fullstack, opID string, fields map[string]oapiparser.FieldConstraint) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for fieldName, fc := range fields {
		d, ok := xdo75FieldDiag(fs, opID, fieldName, fc)
		if ok {
			diags = append(diags, d)
		}
	}
	return diags
}
