//ff:func feature=gen-gogin type=util control=sequence
//ff:what authStoreCallPrelude — otel span start prelude for RefreshRotate/Logout call blocks

package ssac

import "fmt"

// authStoreCallPrelude writes the `callCtx, callSpan := otel.Tracer(...)` line
// (when WrapCalls is true) and returns the accumulated lines plus the name
// of the context variable to use on the call site (`ctx` or `callCtx`).
func (g *methodGen) authStoreCallPrelude(pkgName, callFunc string) ([]string, string) {
	if !g.WrapCalls {
		return nil, "ctx"
	}
	spanName := fmt.Sprintf("call.%s.%s", pkgName, callFunc)
	return []string{
		fmt.Sprintf("callCtx, callSpan := otel.Tracer(\"ssac\").Start(ctx, %q)", spanName),
	}, "callCtx"
}
