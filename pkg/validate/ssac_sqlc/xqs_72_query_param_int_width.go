//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-sqlc
//ff:what XQS-72 — ERROR when OpenAPI query param int width does not match sqlc param int width

package ssac_sqlc

import (
	"fmt"
	"strings"

	"github.com/ettle/strcase"
	"github.com/jinzhu/inflection"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
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

// xqs72CheckSeq checks a single sequence for XQS-72 violations.
func xqs72CheckSeq(
	fn ssac.ServiceFunc,
	seq ssac.Sequence,
	oapiParams map[string]string,
	paramMap map[string]map[string]bool,
	ddlColType map[string]map[string]string,
	queryBodyMap map[string]sqlcparser.QuerySpec,
) []diagnostic.Diagnostic {
	if seq.Type == "call" {
		return nil
	}
	if seq.Model == "" {
		return nil
	}
	queryName := resolveQueryName(seq)
	query, hasQuery := queryBodyMap[queryName]
	if !hasQuery {
		return nil
	}
	sqlcParams, hasSqlc := paramMap[queryName]

	modelName := strings.SplitN(seq.Model, ".", 2)[0]
	tableName := inflection.Plural(strcase.ToSnake(modelName))

	var diags []diagnostic.Diagnostic

	// Check Inputs map (key = PascalCase sqlc param, val = "request.<field>")
	for key, val := range seq.Inputs {
		if d, ok := xqs72CheckEntry(fn, seq, key, val, oapiParams, sqlcParams, hasSqlc, ddlColType, tableName, query.Body); ok {
			diags = append(diags, d)
		}
	}
	return diags
}

// xqs72CheckEntry validates a single Input entry for XQS-72.
func xqs72CheckEntry(
	fn ssac.ServiceFunc,
	seq ssac.Sequence,
	key, val string,
	oapiParams map[string]string,
	sqlcParams map[string]bool,
	hasSqlc bool,
	ddlColType map[string]map[string]string,
	tableName string,
	queryBody string,
) (diagnostic.Diagnostic, bool) {
	if !strings.HasPrefix(val, "request.") {
		return diagnostic.Diagnostic{}, false
	}
	paramName := strings.TrimPrefix(val, "request.")

	// Skip if param not in sqlc params
	if hasSqlc && !sqlcParams[key] {
		return diagnostic.Diagnostic{}, false
	}

	// Skip if key maps to DDL column — XQS-18 handles those
	colName := strcase.ToSnake(key)
	if _, found := xqs18LookupDDLType(ddlColType, tableName, colName); found {
		return diagnostic.Diagnostic{}, false
	}
	if _, found := xqs18LookupDDLType(ddlColType, tableName, paramName); found {
		return diagnostic.Diagnostic{}, false
	}

	// Get OpenAPI format
	oapiFormat, hasOAPI := oapiParams[paramName]
	if !hasOAPI {
		return diagnostic.Diagnostic{}, false
	}
	// Only check int32/int64 formats
	if oapiFormat != "int32" && oapiFormat != "int64" {
		return diagnostic.Diagnostic{}, false
	}

	// Infer sqlc param int width from query body
	sqlcWidth := inferSqlcParamIntWidth(queryBody, strcase.ToSnake(key))
	if sqlcWidth == "" {
		return diagnostic.Diagnostic{}, false
	}

	if oapiFormat == sqlcWidth {
		return diagnostic.Diagnostic{}, false
	}

	return diagnostic.Diagnostic{
		File:  fn.FileName,
		Line:  seq.Line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XQS-72] Input key %q — OpenAPI query param format %s ≠ sqlc param inferred type %s",
			key, oapiFormat, sqlcWidth,
		),
		Advice: fmt.Sprintf(
			"Fix OpenAPI to format: %s or add ::%s cast in the sqlc query",
			sqlcWidth, widthToPGCast(oapiFormat),
		),
	}, true
}

// widthToPGCast maps an OpenAPI int width to the corresponding PG cast token.
func widthToPGCast(width string) string {
	if width == "int64" {
		return "bigint"
	}
	return "int"
}
