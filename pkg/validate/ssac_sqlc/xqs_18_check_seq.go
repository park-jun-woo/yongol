//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-sqlc
//ff:what xqs18CheckSeq — 단일 sequence 의 CRUD Args + Inputs 타입 호환성 검사

package ssac_sqlc

import (
	"strings"

	"github.com/ettle/strcase"
	"github.com/jinzhu/inflection"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xqs18CheckSeq validates a single SSaC Sequence for XQS-18 — covers both
// CRUD Args (source == "request") and the Inputs map (request.* references).
func xqs18CheckSeq(fn ssac.ServiceFunc, seq ssac.Sequence, oapiParams map[string]string, paramMap map[string]map[string]bool, ddlColType map[string]map[string]string) []diagnostic.Diagnostic {
	if seq.Type == "call" {
		return nil
	}
	if seq.Model == "" {
		return nil
	}
	queryName := resolveQueryName(seq)
	sqlcParams, hasSqlc := paramMap[queryName]
	modelName := strings.SplitN(seq.Model, ".", 2)[0]
	tableName := inflection.Plural(strcase.ToSnake(modelName))

	var diags []diagnostic.Diagnostic
	for _, arg := range seq.Args {
		if d, ok := xqs18CheckArg(fn, seq, arg, oapiParams, sqlcParams, hasSqlc, ddlColType, tableName); ok {
			diags = append(diags, d)
		}
	}
	for key, val := range seq.Inputs {
		if d, ok := xqs18CheckInput(fn, seq, key, val, oapiParams, sqlcParams, hasSqlc, ddlColType, tableName); ok {
			diags = append(diags, d)
		}
	}
	return diags
}
