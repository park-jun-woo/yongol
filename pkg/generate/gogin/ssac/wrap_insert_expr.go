//ff:func feature=gen-gogin type=util control=sequence
//ff:what wrapInsertExpr — sqlc INSERT 인자를 InsertExpr 템플릿으로 래핑 (pgtypex bridge + Ptr/non-Ptr 분기 + 정수 캐스트)

package ssac

import (
	"strings"
	"unicode"

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

// isNonPointerInput returns true when the SSaC value resolves to a Go
// value type (not a pointer). Used to select between Ptr and non-Ptr
// pgtypex bridge variants.
//
//   - Literals (true, false, numbers, quoted strings) → non-pointer
//   - request.* path params → non-pointer (oapi-codegen value types)
//   - request.* required body fields → non-pointer
//   - request.* optional body fields → pointer (*string, *int64)
//   - request.* query params → depends on IsRequired
//   - Other expressions (row.Field, var.Field) → assumed pointer-capable
//     (already pgtype), but those are filtered by alreadyPgtype before
//     reaching this function
func (g *methodGen) isNonPointerInput(ssacValue string) bool {
	// Literals are always non-pointer values.
	if isLiteral(ssacValue) {
		return true
	}

	if !strings.HasPrefix(ssacValue, "request.") {
		return false
	}

	field := ssacValue[len("request."):]

	// Path params are always value types in oapi-codegen.
	if g.PathParams[field] {
		return true
	}

	// Query params: required = value type, optional = pointer.
	if qp, isQuery := g.QueryParams[field]; isQuery {
		return qp.IsRequired
	}

	// Body field: required fields are value types, optional are pointers.
	return g.BodyRequiredFields[field]
}

// castIntegerLiteral adds a Go type cast to an integer literal so it
// matches the pgtypex bridge function's expected parameter type.
// e.g. "1" → "int64(1)" for a BIGINT column.
//
// Only applies when rendered is a plain integer literal (digits with
// optional leading minus). Non-integer literals and expressions pass
// through unchanged.
func (g *methodGen) castIntegerLiteral(rendered string, binding types.GoTypeBinding) string {
	if !isIntegerLiteralStr(rendered) {
		return rendered
	}
	goType := goIntTypeFromSqlcType(binding.SqlcGoType)
	if goType == "" {
		return rendered
	}
	return goType + "(" + rendered + ")"
}

// isIntegerLiteralStr returns true when s is a plain integer literal
// (optional leading minus followed by digits only).
func isIntegerLiteralStr(s string) bool {
	if len(s) == 0 {
		return false
	}
	start := 0
	if s[0] == '-' {
		start = 1
	}
	if start >= len(s) {
		return false
	}
	for i := start; i < len(s); i++ {
		if !unicode.IsDigit(rune(s[i])) {
			return false
		}
	}
	return true
}

// goIntTypeFromSqlcType returns the Go integer primitive corresponding to
// a pgtype sqlc type. Returns empty string for non-integer types.
func goIntTypeFromSqlcType(sqlcGoType string) string {
	switch sqlcGoType {
	case "pgtype.Int8":
		return "int64"
	case "pgtype.Int4":
		return "int32"
	case "pgtype.Int2":
		return "int16"
	case "int64":
		return "int64"
	case "int32":
		return "int32"
	default:
		return ""
	}
}
