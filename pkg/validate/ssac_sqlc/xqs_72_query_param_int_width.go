//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-sqlc
//ff:what XQS-72 — OpenAPI query param int width ↔ sqlc param int width 불일치 ERROR

package ssac_sqlc

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xqs72QueryParamIntWidth validates that every SSaC Input whose key does NOT
// map to a DDL column has matching integer widths between OpenAPI (format) and
// the sqlc query body (cast or default).
func xqs72QueryParamIntWidth(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.OpenAPIDoc == nil {
		return nil
	}
	opMap := buildXqs18OperationMap(fs.OpenAPIDoc)
	if len(opMap) == 0 {
		return nil
	}
	ddlColType := buildXqs18DDLColumnTypeMap(fs)
	paramMap := buildQueryParamMap(fs)
	queryBodyMap := buildQueryBodyMap(fs)

	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		op, ok := opMap[fn.Name]
		if !ok {
			continue
		}
		oapiParams := buildXqs18OAPIParamTypeMap(op)
		for _, seq := range fn.Sequences {
			diags = append(diags, xqs72CheckSeq(fn, seq, oapiParams, paramMap, ddlColType, queryBodyMap)...)
		}
	}
	return diags
}
