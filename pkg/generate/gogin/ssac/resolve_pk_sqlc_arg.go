//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what resolvePKSqlcArg — PK 컬럼 바인딩으로 sqlcArg 표현식·pgtypex import 산출

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// resolvePKSqlcArg wraps ridExpr with InsertExpr from the PK column binding
// and collects pgtypex-related imports. Returns the (possibly wrapped)
// expression and any additional imports.
func resolvePKSqlcArg(pkCol *ddl.Column, ridExpr string) (string, []string) {
	if pkCol == nil {
		return ridExpr, nil
	}
	binding := types.MapPGType(*pkCol)
	if binding.InsertExpr == "" || binding.InsertExpr == "{var}" {
		return ridExpr, nil
	}
	sqlcArg := types.Expand(binding.InsertExpr, "", "", ridExpr)
	var imports []string
	for _, imp := range binding.Imports {
		if strings.Contains(imp, "pgtypex") {
			imports = append(imports, `"`+imp+`"`)
		}
	}
	return sqlcArg, imports
}
