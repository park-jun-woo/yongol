//ff:func feature=gen-gogin type=util control=sequence
//ff:what resolvePgtypeFieldExpr — dotted 필드 접근(var.Field)이 pgtype 컬럼이면 pgtypex 변환 표현식 반환

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
)

// resolvePgtypeFieldExpr checks whether a dotted expression like "user.Name"
// refers to a pgtype column on the variable's model type. If so it returns
// the pgtypex conversion expression and the required import paths; otherwise
// it returns ("", nil).
//
// The resolution works by:
//  1. Splitting the expression on "." to extract the variable name and field.
//  2. Looking up the variable's model type via VarTypes (populated from SSaC
//     result bindings).
//  3. Resolving the DDL column for the (modelName, fieldName) pair.
//  4. Checking if the column maps to a KindPgtype binding.
//  5. If so, returning the ConvertExpr expanded with the variable name and
//     field name, along with the binding's Imports.
func (g *methodGen) resolvePgtypeFieldExpr(varExpr string) (string, []string) {
	parts := strings.SplitN(varExpr, ".", 2)
	if len(parts) != 2 {
		return "", nil
	}
	varName := parts[0]
	fieldName := parts[1]

	modelName, ok := g.VarTypes[varName]
	if !ok {
		return "", nil
	}
	// Trim slice prefix — VarTypes may store "[]Workflow" for list results.
	modelName = strings.TrimPrefix(modelName, "[]")

	// Pass the PascalCase sqlc field name directly. lookupDDLColumn applies
	// caseconv.PascalToSnake internally, which handles acronyms correctly
	// (ID→id, OrgID→org_id, URL→url). A prior re-lowercasing step produced
	// broken keys for leading/inner acronyms (ID→"iD"→"i_d" miss).
	col := lookupDDLColumn(g.DDLTables, modelName, fieldName)
	if col == nil {
		return "", nil
	}
	binding := types.MapPGType(*col)
	if binding.Kind != types.KindPgtype {
		return "", nil
	}
	// ConvertExpr uses {row}.{field} placeholders. For a dotted access
	// like "user.Name", {row}=varName and {field}=fieldName (PascalCase
	// as sqlc emits it).
	expr := types.Expand(binding.ConvertExpr, varName, fieldName, "")

	// Collect only the imports that the handler file actually references.
	// The pgtypex bridge functions are called directly; the pgtype import
	// is needed only by the sqlc-generated db package, not by the handler.
	// Other imports (e.g. "time" for FromPgTimestamptz) are included.
	var quoted []string
	for _, imp := range binding.Imports {
		// pgtype import is used by db package, not by the handler code.
		// runtime/types is used by the convert-func file (openapi_types
		// alias) but never by the dotted-access conversion handler, which
		// assigns the conversion expression directly without naming the type.
		if imp == "github.com/jackc/pgx/v5/pgtype" ||
			imp == "github.com/oapi-codegen/runtime/types" {
			continue
		}
		quoted = append(quoted, `"`+imp+`"`)
	}
	return expr, quoted
}
