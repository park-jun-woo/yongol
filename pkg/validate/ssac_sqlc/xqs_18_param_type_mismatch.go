//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-sqlc
//ff:what XQS-18 — ERROR when a SSaC Input request.* OpenAPI param type does not match the sqlc/DDL type

package ssac_sqlc

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// openAPITypeCompatible reports whether an OpenAPI type is compatible with a Go DDL type.
// Format-specific keys (int64, int32, …) require exact match.
// Base-type keys (integer, number, …) allow loose match when format is absent.
var openAPITypeCompatible = map[string]map[string]bool{
	"int64":   {"int64": true},
	"int32":   {"int32": true, "int": true},
	"int16":   {"int16": true},
	"uuid":    {"pgtype.UUID": true, "string": true},
	"float32": {"float32": true},
	"float64": {"float64": true},

	"integer": {"int32": true, "int64": true, "int16": true, "int": true},
	"number":  {"float32": true, "float64": true},
	"string":  {"string": true},
	"boolean": {"bool": true},
}

// xqs18ParamTypeMismatch validates XQS-18: when a SSaC Input references
// request.* and the same key exists in sqlc Params, the OpenAPI param type
// must be compatible with the DDL column Go type.
// Skip: seq.Type == "call".
func xqs18ParamTypeMismatch(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.OpenAPIDoc == nil {
		return nil
	}
	opMap := buildXqs18OperationMap(fs.OpenAPIDoc)
	if len(opMap) == 0 {
		return nil
	}
	ddlColType := buildXqs18DDLColumnTypeMap(fs)
	paramMap := buildQueryParamMap(fs)

	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		op, ok := opMap[fn.Name]
		if !ok {
			continue
		}
		oapiParams := buildXqs18OAPIParamTypeMap(op)
		for _, seq := range fn.Sequences {
			diags = append(diags, xqs18CheckSeq(fn, seq, oapiParams, paramMap, ddlColType)...)
		}
	}
	return diags
}
