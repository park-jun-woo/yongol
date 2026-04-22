//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildCall — @call 시퀀스 빌더 (ssac 내장 vs 프로젝트 custom 분기)

package ssac

import (
	"fmt"
	"strings"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// ssacBuiltinPkgs are packages provided by github.com/park-jun-woo/ssac/pkg/.
//
// `auth` became fully ssac-origin in Phase003 — JWT Issue/Refresh/Verify and
// the refresh rotation runtime all live in ssac/pkg/auth. SSaC @call targets
// still resolve through the project-local `internal/auth` alias surface
// because build_call emits `auth.Foo(...)` with no package-qualified import;
// the project's reexport.go re-alias maps those symbols back to ssac/pkg/auth.
// That is why `auth` is NOT listed here: callers should import
// `modulePath/internal/auth`, which is the else-branch in buildCall.
var ssacBuiltinPkgs = map[string]bool{
	"authz": true, "cache": true, "config": true,
	"crypto": true, "file": true, "image": true, "mail": true,
	"pagination": true, "queue": true, "session": true, "storage": true, "text": true,
}

// ssacCtxFirstPkgs lists ssac/pkg packages whose public functions accept
// context.Context as the first argument (Phase004b migration). buildCall
// injects `ctx, ` before the Request literal when the target package is in
// this set. Packages not listed here retain the legacy `func X(req)` shape.
//
// `auth` is intentionally absent — it is dual-origin (ssac + project-local)
// and bcrypt/JWT operations are CPU-bound, so ctx adds no cancellation
// value; migration there requires coordinated changes to
// `internal/auth/reexport.go` and project-local IssueToken/VerifyToken.
var ssacCtxFirstPkgs = map[string]bool{
	"mail":    true,
	"storage": true,
	"cache":   true,
	"session": true,
	"file":    true,
}

func (g *methodGen) buildCall(seq ssacparser.Sequence) ([]string, []string) {
	parts := strings.SplitN(seq.Model, ".", 2)
	pkgName := parts[0]
	callFunc := parts[1]
	fields := g.mapFields(seq.Inputs)

	// Phase003 — auth.IssueToken / auth.RefreshToken now accept a single
	// `Claims any` field. SSaC @call authors write natural claim fields
	// (ID, Email, Role, OrgID); wrap them in auth.Claim{...} so the
	// generated call matches the ssac/pkg/auth signature without forcing
	// SSaC rewrites.
	if pkgName == "auth" && (callFunc == "IssueToken" || callFunc == "RefreshToken") {
		fields = "Claims: auth.Claim{" + fields + "}"
	}

	// Phase009 — auth.RefreshRotate and auth.Logout depart from the SSaC
	// "single XRequest struct" convention. Their Go signatures take
	// (ctx, *RefreshStore, refreshToken string) because the rotation
	// runtime needs the store handle to Consume+Create atomically and the
	// SSaC authors shouldn't thread the store through @call literals.
	// Dispatch to a dedicated emitter so the rest of buildCall's
	// span-wrap / error-mapping / import logic stays reusable.
	if pkgName == "auth" && (callFunc == "RefreshRotate" || callFunc == "Logout") {
		return g.buildAuthRefreshStoreCall(seq, callFunc)
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
	ctxPrefix := ""
	if ssacCtxFirstPkgs[pkgName] {
		ctxPrefix = "ctx, "
	}

	var lines []string
	// Phase009 optional span wrap: when manifest's
	// observability.tracing.wrap_calls is true, emit an explicit child span
	// for every @call so the trace tree shows each runtime-package call as
	// its own timing node. A closure pattern keeps the span tight around
	// the call itself — End() runs as soon as the returned ctx has been
	// consumed so later sequences don't accidentally inherit the wrapper
	// ctx, and an early return from the handler still flushes the span.
	spanCtxVar := "ctx"
	if g.WrapCalls {
		spanName := fmt.Sprintf("call.%s.%s", pkgName, callFunc)
		lines = append(lines,
			fmt.Sprintf("callCtx, callSpan := otel.Tracer(\"ssac\").Start(ctx, %q)", spanName),
		)
		spanCtxVar = "callCtx"
	}

	// Swap the ctx identifier passed into the runtime call when wrap_calls
	// is active so the span is the effective parent for that invocation.
	if ctxPrefix != "" {
		ctxPrefix = spanCtxVar + ", "
	}

	lines = append(lines,
		fmt.Sprintf("%s, err %s %s.%s(%s%s.%sRequest{%s})", varName, assign, pkgName, callFunc, ctxPrefix, pkgName, callFunc, fields),
	)
	if g.WrapCalls {
		// End the span before the error branch so the span records the
		// actual call duration (not the slog / return overhead).
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

	// Phase004 — when the @call target is auth.RefreshToken, emit a follow-up
	// server.RefreshStore.Create(...) so the newly minted refresh JWT is
	// persisted in one shot. SSaC authors only write `@call auth.RefreshToken`;
	// the DB Create is generator-synthesized so user-authored SSaC stays
	// claim-agnostic and the Server struct's RefreshStore field (added by
	// generate_server_go.go) becomes the single injection point. Assumes a
	// variable binding exists — without `refresh = ...` there is nothing to
	// persist — and that Inputs match the Claim struct fields emitted above.
	if pkgName == "auth" && callFunc == "RefreshToken" && varName != "_" && !g.IsSubscribe {
		claimLit := "auth.Claim{" + g.mapFields(seq.Inputs) + "}"
		lines = append(lines,
			fmt.Sprintf("if err := server.RefreshStore.Create(ctx, %s.RefreshToken, %s, %s.ExpiresAt); err != nil {", varName, claimLit, varName),
			fmt.Sprintf("\tslog.Error(\"handler: 5xx\", \"op\", %q, \"status\", 500, \"err\", err)", g.FuncName),
			fmt.Sprintf("\treturn api.%s500JSONResponse{Error: %q, Code: strPtr(%q)}, nil", g.FuncName, neutralMessage(500), neutralCode(500)),
			"}",
		)
	}

	// Phase020 — after auth.RefreshToken (which completes the IssueToken +
	// RefreshToken + Create sequence typical of a Login handler), emit a
	// call to auth.SetAuthCookies so the generated handler ships Set-Cookie
	// headers matching the JSON body. The call is a no-op at runtime when
	// BACKEND_AUTH_MODE=bearer, so bearer deployments pay only the function
	// call cost (the cookie-mode conditional lives inside auth.Configure,
	// not in every generated handler).
	//
	// The paired access token is bound to the variable produced by the
	// immediately-preceding @call auth.IssueToken (typical SSaC pattern:
	// `token = auth.IssueToken(...)`). We detect that by scanning the
	// method's accumulated token variable name — set below on IssueToken.
	if pkgName == "auth" && callFunc == "RefreshToken" && varName != "_" && !g.IsSubscribe && g.AccessTokenVar != "" {
		lines = append(lines,
			fmt.Sprintf("// Phase020 — emit Set-Cookie for access+refresh. Runtime no-op when Mode==\"bearer\"."),
			fmt.Sprintf("if ginCtx, ok := ctx.(*gin.Context); ok {"),
			fmt.Sprintf("\tauth.SetAuthCookies(ginCtx, %s.AccessToken, %s.RefreshToken)", g.AccessTokenVar, varName),
			"}",
		)
	}
	if pkgName == "auth" && callFunc == "IssueToken" && varName != "_" && !g.IsSubscribe {
		// Record the access-token variable so the RefreshToken branch
		// above can pair it into auth.SetAuthCookies. Only the most
		// recent IssueToken wins — projects chain at most one in a
		// given handler per Phase020 convention (Login).
		g.AccessTokenVar = varName
	}

	var imps []string
	if ssacBuiltinPkgs[pkgName] {
		imps = append(imps, fmt.Sprintf(`"github.com/park-jun-woo/ssac/pkg/%s"`, pkgName))
	} else {
		imps = append(imps, fmt.Sprintf(`"%s/internal/%s"`, g.ModulePath, pkgName))
	}
	if g.IsSubscribe {
		imps = append(imps, `"fmt"`)
	} else {
		imps = append(imps, `"log/slog"`)
	}
	if g.WrapCalls {
		imps = append(imps,
			`"go.opentelemetry.io/otel"`,
		)
	}
	// Phase020 — gin.Context import needed for the SetAuthCookies emission
	// (only on auth.RefreshToken follow-up). Safe to always add on auth.*;
	// deduplicate_imports will drop unused entries in other call paths.
	if pkgName == "auth" && (callFunc == "RefreshToken" || callFunc == "IssueToken") {
		imps = append(imps, `"github.com/gin-gonic/gin"`)
	}
	return lines, imps
}
