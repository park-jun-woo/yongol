//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-sqlc
//ff:what XQS-19 — SSaC 가 호출하는 DB-using ssac built-in 에 대응 sqlc 쿼리 존재 강제

package ssac_sqlc

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xqs19SsacBuiltinQueryRequired validates XQS-19: when a SSaC function
// invokes a DB-using ssac built-in (`@call cache.Set`, `@publish topic`,
// `@subscribe topic`), every sqlc query declared as `used_by:` for that
// method in the package's interface.yaml must exist in fs.SQLcQueries.
func xqs19SsacBuiltinQueryRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || len(fs.SsacInterfaces) == 0 {
		return nil
	}
	have := collectHaveQueries(fs)
	var diags []diagnostic.Diagnostic
	for _, f := range fs.ServiceFuncs {
		diags = append(diags, collectFuncCallDiags(f, fs.SsacInterfaces, have)...)
		diags = append(diags, collectFuncSubscribeDiags(f, fs.SsacInterfaces, have)...)
	}
	return diags
}

// _ pins a reference to ssacmeta.Port so compilers flag an interface.yaml
// schema change (e.g. UsedBy rename) at the earliest possible test run.
var _ = ssacmeta.Port{}
