//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildAuthRefreshStoreCall — auth.RefreshRotate / auth.Logout 전용 emitter (Phase009)

package ssac

import (
	"fmt"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// buildAuthRefreshStoreCall emits the call block for the two Phase009 auth
// endpoints that do NOT follow the SSaC "one XRequest struct" calling
// convention: auth.RefreshRotate and auth.Logout. Both take
// (ctx, *RefreshStore, refreshToken string). The emitter:
//
//   - Pulls the single "RefreshToken" input from seq.Inputs and maps its
//     right-hand side via g.mapValue (so `request.refresh_token` becomes
//     `request.Body.RefreshToken`, etc).
//   - Emits `server.RefreshStore` as the store argument — the Server struct
//     gained a RefreshStore pointer in Phase004 and block_auth_init wires
//     it at boot.
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

	lines = append(lines,
		fmt.Sprintf("%s, err %s %s.%s(%s, server.RefreshStore, %s)",
			varName, assign, pkgName, callFunc, spanCtxVar, tokenArg),
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
	// auth package is project-local (internal/auth reexport).
	imps = append(imps, fmt.Sprintf(`"%s/internal/%s"`, g.ModulePath, pkgName))
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
