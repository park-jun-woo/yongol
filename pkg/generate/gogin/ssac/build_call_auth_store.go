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
// The handler therefore passes `nil` instead of `server.RefreshStore`,
// which is no longer a field on the Server struct.
//
// detectReuseLogoutAll (bool) is a hard-coded literal — Phase002 hoists it
// out of the removed `&auth.RefreshStore{DetectReuseLogoutAll: true}`
// literal in main.go and surfaces it as the fourth argument of
// RefreshRotate. Default false keeps zenflow behavior; future manifest
// config can flip the literal.
//
// The emitter:
//
//   - Pulls the single "RefreshToken" input from seq.Inputs and maps its
//     right-hand side via g.mapValue.
//   - Emits `nil` as the store argument so ssac falls back to its
//     package-level singleton.
//   - Reuses the same tracing span-wrap + error-to-JSON branch as buildCall
//     so generated handlers look identical except for the call shape.
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

	var lines []string
	spanCtxVar := "ctx"
	if g.WrapCalls {
		spanName := fmt.Sprintf("call.%s.%s", pkgName, callFunc)
		lines = append(lines,
			fmt.Sprintf("callCtx, callSpan := otel.Tracer(\"ssac\").Start(ctx, %q)", spanName),
		)
		spanCtxVar = "callCtx"
	}

	// Phase002 (ssac/purify) — store is passed as nil so ssac falls back to
	// the auth.Init(...) singleton installed in main.go. RefreshRotate now
	// also accepts a detectReuseLogoutAll bool; Logout keeps the 3-arg shape.
	var callExpr string
	if callFunc == "RefreshRotate" {
		callExpr = fmt.Sprintf("%s.%s(%s, nil, %s, false)", pkgName, callFunc, spanCtxVar, tokenArg)
	} else {
		callExpr = fmt.Sprintf("%s.%s(%s, nil, %s)", pkgName, callFunc, spanCtxVar, tokenArg)
	}
	lines = append(lines,
		fmt.Sprintf("%s, err %s %s", varName, assign, callExpr),
	)
	if g.WrapCalls {
		lines = append(lines, "callSpan.End()")
	}
	lines = append(lines, "if err != nil {")
	if g.IsSubscribe {
		lines = append(lines, fmt.Sprintf("\treturn fmt.Errorf(\"%s.%s: %%w\", err)", pkgName, callFunc))
	} else {
		lines = append(lines,
			fmt.Sprintf("\t%s(\"handler: %s\", \"op\", %q, \"status\", %d, \"err\", err)", logLevelFuncForStatus(status), logTagForStatus(status), g.FuncName, status),
			fmt.Sprintf("\treturn api.%s%dJSONResponse{Error: %q, Code: strPtr(%q)}, nil", g.FuncName, status, msg, neutralCode(status)),
		)
	}
	lines = append(lines, "}")

	// Phase020 — RefreshRotate produces both a new access and a new refresh
	// token; emit Set-Cookie for both. Logout invalidates the refresh row;
	// emit Max-Age=0 Set-Cookie headers to clear the browser's copy so the
	// next request genuinely lands unauthenticated instead of replaying the
	// (now-revoked) cookie. Both helpers are Mode-gated at runtime: bearer
	// deployments see a no-op.
	if !g.IsSubscribe && varName != "_" && callFunc == "RefreshRotate" {
		lines = append(lines,
			"// Phase020 — refresh rotation emits fresh Set-Cookie headers.",
			"if ginCtx, ok := ctx.(*gin.Context); ok {",
			fmt.Sprintf("\tauth.SetAuthCookies(ginCtx, %s.AccessToken, %s.RefreshToken)", varName, varName),
			"}",
		)
	}
	if !g.IsSubscribe && callFunc == "Logout" {
		lines = append(lines,
			"// Phase020 — Logout clears the session cookies (Max-Age=0).",
			"if ginCtx, ok := ctx.(*gin.Context); ok {",
			"\tauth.ClearAuthCookies(ginCtx)",
			"}",
		)
	}

	var imps []string
	// Phase001 UserClaimUnification — `auth` is back on ssac/pkg/auth for
	// all emission paths; RefreshRotate/Logout live in ssac/pkg/auth.
	imps = append(imps, fmt.Sprintf(`"github.com/park-jun-woo/ssac/pkg/%s"`, pkgName))
	if g.IsSubscribe {
		imps = append(imps, `"fmt"`)
	} else {
		imps = append(imps, `"log/slog"`)
	}
	if g.WrapCalls {
		imps = append(imps, `"go.opentelemetry.io/otel"`)
	}
	// Phase020 — gin import needed for the ctx.(*gin.Context) assertion
	// above. Included on both RefreshRotate and Logout; subscribe path
	// skips the emission so no import is needed.
	if !g.IsSubscribe && (callFunc == "RefreshRotate" || callFunc == "Logout") {
		imps = append(imps, `"github.com/gin-gonic/gin"`)
	}
	return lines, imps
}
