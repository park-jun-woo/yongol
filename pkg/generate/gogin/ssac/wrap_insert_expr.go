//ff:func feature=gen-gogin type=util control=sequence
//ff:what wrapInsertExpr — sqlc INSERT 인자를 InsertExpr 템플릿으로 래핑 (pgtypex bridge)

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
// Returns the wrapped expression and any additional imports needed
// (quoted for writeMethodFile format).
func (g *methodGen) wrapInsertExpr(inputKey, rendered string) (string, []string) {
	col := g.lookupSQLCMethodColumn(inputKey)
	if col == nil {
		return rendered, nil
	}
	binding := types.MapPGType(*col)
	if binding.InsertExpr == "" || binding.InsertExpr == "{var}" {
		return rendered, nil
	}
	wrapped := types.Expand(binding.InsertExpr, "", "", rendered)
	// binding.Imports holds bare paths; handler emit expects quoted form.
	var imports []string
	for _, imp := range binding.Imports {
		if strings.Contains(imp, "pgtypex") {
			imports = append(imports, `"`+imp+`"`)
		}
	}
	return wrapped, imports
}
