//ff:func feature=validate type=util control=sequence topic=func-check
//ff:what xfs39Advice — 누락된 @call 대상에 대한 contextual advice 생성

package ssac_func

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xfs39Advice returns contextual advice for a missing @call target. When the
// call targets a builtin ssac package, it lists the available functions in
// that package (merged from YongolPkgSpecs and Ground). Otherwise it falls
// back to the generic "define under pkg/" message.
func xfs39Advice(model string, fs *yongol.Fullstack) string {
	idx := strings.IndexByte(model, '.')
	if idx <= 0 {
		return "Define function " + model + " under pkg/"
	}
	pkg := model[:idx]
	if !builtinPackages[pkg] {
		return "Define function " + model + " under pkg/"
	}
	names := collectBuiltinFuncNames(pkg, fs)
	if len(names) == 0 {
		return "Package " + pkg + " is a builtin ssac package but no functions were loaded. Check ssac/pkg installation."
	}
	return "Available " + pkg + " functions: " + strings.Join(names, ", ")
}
