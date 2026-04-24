//ff:func feature=validate type=rule control=iteration dimension=1 topic=states
//ff:what XSM-27 — stateful POST/PUT/DELETE 의 @state / @state-neutral 의도 선언 강제

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xsm27StateIntentDeclaration validates XSM-27: any POST/PUT/DELETE that
// targets a stateful resource (via a `{id}` path parameter) and reads that
// resource with `@get <Model>.FindByID({ID: request.id})` must either declare
// a `@state` guard or explicitly opt out with `// @state-neutral`.
//
// Fires WARNING when all the following hold:
//  1. The OpenAPI operation's path contains an `{id}`-ish parameter
//  2. The HTTP method is POST, PUT, or DELETE
//  3. The first path segment maps to a stateful resource (state diagram
//     exists and DDL DEFAULT matches the diagram's initial state — XDM-28 linkage)
//  4. The SSaC function reads that resource via `@get <StatefulModel>.FindByID({ID: request.id})`
//  5. Neither a `@state` sequence nor a `// @state-neutral` annotation is present
func xsm27StateIntentDeclaration(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Paths == nil {
		return nil
	}
	g := fs.Ground()
	if g == nil {
		return nil
	}
	funcByName := buildFuncByName(fs.ServiceFuncs)
	var diags []diagnostic.Diagnostic
	for pathStr, item := range fs.OpenAPIDoc.Paths.Map() {
		diags = append(diags, xsm27DiagsForPath(pathStr, item, fs.StateDiagrams, g, funcByName)...)
	}
	return diags
}
