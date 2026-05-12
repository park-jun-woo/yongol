//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-sqlc
//ff:what xqs72CheckSeq — 단일 시퀀스의 Input 항목에 대해 XQS-72 위반 검사

package ssac_sqlc

import (
	"strings"

	"github.com/ettle/strcase"
	"github.com/jinzhu/inflection"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xqs72CheckSeq checks a single sequence for XQS-72 violations.
func xqs72CheckSeq(
	fn ssac.ServiceFunc,
	seq ssac.Sequence,
	oapiParams map[string]string,
	paramMap map[string]map[string]bool,
	ddlColType map[string]map[string]string,
	queryBodyMap map[string]sqlcparser.QuerySpec,
) []diagnostic.Diagnostic {
	if seq.Type == "call" || seq.Model == "" {
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
	for key, val := range seq.Inputs {
		if d, ok := xqs72CheckEntry(fn, seq, key, val, oapiParams, sqlcParams, hasSqlc, ddlColType, tableName, query.Body); ok {
			diags = append(diags, d)
		}
	}
	return diags
}
