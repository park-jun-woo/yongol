//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what methodGen.sqlcArgsMulti — 입력이 2개 이상일 때 sqlc Params 구조체 리터럴 생성 (JSONB 리터럴 wrap + nullable zero-fill 포함)
package ssac

import (
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
)

// Validate (XQS-14/16) guarantees key matches sqlc field name. Use as-is.
func (g *methodGen) sqlcArgsMulti(method string, inputs map[string]string) (preamble []string, args string, imports []string) {
	var fields []string
	for k, v := range inputs {
		raw, pre, needsJSON := g.maybeMarshalJSONB(k, v)
		if needsJSON {
			preamble = append(preamble, pre...)
			imports = append(imports, `"encoding/json"`)
			fields = append(fields, k+": "+raw)
			continue
		}
		rendered := g.mapValue(v)
		rendered = g.wrapJSONBLiteral(k, rendered)
		alreadyPgtype := !strings.HasPrefix(v, "request.") && !isLiteral(v)
		rendered, extraImports := g.wrapInsertExpr(k, rendered, alreadyPgtype, v)
		imports = append(imports, extraImports...)
		fields = append(fields, k+": "+rendered)
	}

	// BUG-072 Pattern 1: fill missing nullable sqlc params with pgtype zero
	// values so the struct literal compiles (nil is not assignable to pgtype
	// structs).
	zeroFields, zeroImports := g.fillMissingNullableParams(method, inputs)
	fields = append(fields, zeroFields...)
	imports = append(imports, zeroImports...)

	sort.Strings(fields)
	return preamble, "ctx, db." + method + "Params{" + strings.Join(fields, ", ") + "}", imports
}

// fillMissingNullableParams finds sqlc Params fields that the SSaC Inputs
// do not provide and, when the underlying DDL column is nullable (pgtype
// family), emits a zero-value literal (e.g. "pgtype.Int8{}"). This
// prevents Go compile errors where the struct literal would otherwise
// leave a pgtype field at its Go zero value via omission — which is
// actually valid Go, but when an upstream transform inserts explicit nil
// for missing fields the result is "cannot use nil as pgtype.XXX value".
// Emitting the explicit zero value is always safe: pgtype zero = SQL NULL.
func (g *methodGen) fillMissingNullableParams(method string, inputs map[string]string) ([]string, []string) {
	// Find the QuerySpec matching this method name.
	var params []string
	for _, q := range g.SQLcQueries {
		if q.Name == method {
			params = q.Params
			break
		}
	}
	if len(params) == 0 {
		return nil, nil
	}

	model := g.modelForSQLCMethod(method)
	if model == "" {
		return nil, nil
	}

	var fields []string
	var imports []string
	needPgtype := false

	for _, param := range params {
		if _, provided := inputs[param]; provided {
			continue
		}
		col := lookupDDLColumn(g.DDLTables, model, param)
		if col == nil {
			continue
		}
		binding := types.MapPGType(*col)
		if binding.Kind != types.KindPgtype {
			continue
		}
		// Emit explicit zero value: pgtype.Int8{}, pgtype.Text{}, etc.
		fields = append(fields, param+": "+binding.SqlcGoType+"{}")
		needPgtype = true
	}

	if needPgtype {
		imports = append(imports, `"github.com/jackc/pgx/v5/pgtype"`)
	}
	return fields, imports
}
