//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what lookupEvalSpec — @eval pkg.Pascal → FuncSpec (project + yongol pkg)

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

// lookupEvalSpec resolves an @eval Model (e.g. "billing.IsZeroBalance") to its
// FuncSpec. ProjectFuncSpecs override YongolPkgSpecs. Returns nil when no spec
// matches.
func lookupEvalSpec(model string, project, builtin []funcspec.FuncSpec) *funcspec.FuncSpec {
	pkgName, methodPascal := splitEvalModel(model)
	if pkgName == "" || methodPascal == "" {
		return nil
	}
	camel := lcFirstEval(methodPascal)
	for i := range project {
		if project[i].Package == pkgName && project[i].Name == camel {
			return &project[i]
		}
	}
	for i := range builtin {
		if builtin[i].Package == pkgName && builtin[i].Name == camel {
			return &builtin[i]
		}
	}
	return nil
}
