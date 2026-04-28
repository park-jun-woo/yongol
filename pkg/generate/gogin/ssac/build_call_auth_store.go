//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildAuthRefreshStoreCall — auth.RefreshRotate / auth.Logout 전용 emitter (Phase009)

package ssac

import (
	"fmt"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// buildAuthRefreshStoreCall emits the call block for the two Phase009 auth
// endpoints that do NOT follow the SSaC "one XRequest struct" calling
// convention: auth.RefreshRotate and auth.Logout.
//
// Phase002 (ssac/purify) updates the call shape:
//
//   - auth.RefreshRotate(ctx, store, refreshToken, detectReuseLogoutAll bool)
//   - auth.Logout(ctx, store, refreshToken)
//
// with `store` now a nil-fallback: when nil, ssac/pkg/auth reads the
// singleton installed at boot by auth.Init(infraauth.NewPostgres(queries)).
func (g *methodGen) buildAuthRefreshStoreCall(seq ssacparser.Sequence, callFunc string) ([]string, []string) {
	pkgName := "auth"
	tokenArg := g.mapValue(seq.Inputs["RefreshToken"])
	if tokenArg == "" {
		// Defensive: without a RefreshToken input we can't emit a valid
		// call. Fall back to the raw literal so the Go compiler surfaces
		// the error — preferable to silently producing "".
		tokenArg = `""`
	}

	status := g.resolveErrCode(resolveCallErrStatus(seq.ErrStatus, pkgName, callFunc, g.ProjectFuncs, g.BuiltinFuncs))

	varName := "_"
	if seq.Result != nil {
		varName = seq.Result.Var
	}
	assign := g.assignOp(varName != "_")

	msg := seq.Message
	if msg == "" {
		msg = neutralMessage(status)
	}

	lines, spanCtxVar := g.authStoreCallPrelude(pkgName, callFunc)
	callExpr := authStoreCallExpr(pkgName, callFunc, spanCtxVar, tokenArg)
	lines = append(lines, fmt.Sprintf("%s, err %s %s", varName, assign, callExpr))
	if g.WrapCalls {
		lines = append(lines, "callSpan.End()")
	}
	lines = append(lines, g.authStoreErrBranch(pkgName, callFunc, status, msg)...)
	lines = append(lines, g.authStoreCookieEmit(callFunc, varName)...)

	return lines, g.authStoreImports(pkgName, callFunc)
}
