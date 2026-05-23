//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what wrapInsertExpr — sqlc INSERT 인자를 InsertExpr 템플릿으로 래핑 (pgtypex bridge + Ptr/non-Ptr 분기 + 정수 캐스트)

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
)

// wrapInsertExpr applies the column's InsertExpr template to the rendered
// value expression. When the column binding's InsertExpr is a simple
// passthrough ("{var}"), returns rendered unchanged. When it contains a
// bridge call (e.g. "pgtypex.ToPgUUID({var})"), expands the template.
//
// alreadyPgtype signals that the rendered value is already a pgtype struct
// (e.g. a row field from a previous sqlc call). In that case InsertExpr is
// skipped to avoid double-wrapping.
//
// ssacValue is the original SSaC input value (e.g. "request.recorded_at",
// "false", "1"). Used to determine pointer-ness (BUG-072 Pattern 2) and
// to detect integer literals requiring Go type casts (BUG-072 Pattern 3).
//
// Returns the wrapped expression and any additional imports needed
// (quoted for writeMethodFile format).
func (g *methodGen) wrapInsertExpr(inputKey, rendered string, alreadyPgtype bool, ssacValue string) (string, []string) {
	if alreadyPgtype {
		return rendered, nil
	}
	col := g.lookupSQLCMethodColumn(inputKey)
	if col == nil {
		return rendered, nil
	}
	binding := types.MapPGType(*col)
	if binding.InsertExpr == "" || binding.InsertExpr == "{var}" {
		return rendered, nil
	}

	// BUG-072 Pattern 1b: nil literal cannot be passed to a pgtypex bridge
	// function (e.g. ToPgInt8(nil) — nil is not assignable to int64).
	// Emit the pgtype zero value directly: pgtype.Int8{}, pgtype.Text{}, etc.
	if rendered == "nil" && binding.Kind == types.KindPgtype {
		return binding.SqlcGoType + "{}", []string{`"github.com/jackc/pgx/v5/pgtype"`}
	}

	// BUG-072 Pattern 3: integer literals need an explicit Go type cast so
	// the pgtypex bridge function receives the correct sized integer.
	// e.g. "1" → "int64(1)" when the column is BIGINT (pgtype.Int8).
	rendered = g.castIntegerLiteral(rendered, binding)

	expr := binding.InsertExpr

	// BUG-072 Pattern 2: pgtypex ToPgXxxPtr variants expect a pointer
	// (*string, *int64, …). When the SSaC input resolves to a non-pointer
	// Go expression (required body field, path param, or literal), switch
	// to the non-Ptr variant (ToPgXxx) which accepts a value.
	if strings.Contains(expr, "Ptr(") && g.isNonPointerInput(ssacValue) {
		expr = strings.Replace(expr, "Ptr(", "(", 1)
	}

	wrapped := types.Expand(expr, "", "", rendered)
	// binding.Imports holds bare paths; handler emit expects quoted form.
	var imports []string
	for _, imp := range binding.Imports {
		if strings.Contains(imp, "pgtypex") {
			imports = append(imports, `"`+imp+`"`)
		}
	}
	return wrapped, imports
}
