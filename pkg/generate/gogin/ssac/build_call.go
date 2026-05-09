//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildCall — @call 시퀀스 빌더 (ssac 내장 vs 프로젝트 custom 분기)

package ssac

import (
	"fmt"
	"strings"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// ssacCtxFirstPkgs lists ssac/pkg packages whose public functions accept
// context.Context as the first argument (Phase004b migration). buildCall
// injects `ctx, ` before the Request literal when the target package is in
// this set. Packages not listed here retain the legacy `func X(req)` shape.
//
// `auth` is intentionally absent — bcrypt/JWT operations are CPU-bound, so
// ctx adds no cancellation value; keeping auth on the legacy `func X(req)`
// shape matches the actual ssac/pkg/auth surface.
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
	fields := wrapAuthClaimsFields(pkgName, callFunc, g.mapFields(seq.Inputs))

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
	spanCtxVar := "ctx"
	if g.WrapCalls {
		lines = append(lines, buildCallSpanOpenLines(pkgName, callFunc)...)
		spanCtxVar = "callCtx"
	}
	if ctxPrefix != "" {
		ctxPrefix = spanCtxVar + ", "
	}

	lines = append(lines,
		fmt.Sprintf("%s, err %s %s.%s(%s%s.%sRequest{%s})", varName, assign, pkgName, callFunc, ctxPrefix, pkgName, callFunc, fields),
	)
	if g.WrapCalls {
		lines = append(lines, "callSpan.End()")
	}
	lines = append(lines, g.buildCallErrorLines(pkgName, callFunc, msg, status)...)

	if pkgName == "auth" && callFunc == "RefreshToken" && varName != "_" && !g.IsSubscribe {
		lines = append(lines, g.buildCallRefreshCreateLines(seq, varName)...)
	}
	if pkgName == "auth" && callFunc == "RefreshToken" && varName != "_" && !g.IsSubscribe && g.AccessTokenVar != "" {
		lines = append(lines, g.buildCallSetAuthCookiesLines(varName)...)
	}
	if pkgName == "auth" && callFunc == "IssueToken" && varName != "_" && !g.IsSubscribe {
		g.AccessTokenVar = varName
	}

	return lines, g.buildCallImports(pkgName, callFunc, varName)
}
