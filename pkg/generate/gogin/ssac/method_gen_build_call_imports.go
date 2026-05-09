//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildCallImports — @call emission 시 필요한 import 경로 수집

package ssac

import "fmt"

func (g *methodGen) buildCallImports(pkgName, callFunc, varName string) []string {
	var imps []string
	if fullPath, ok := g.ImportMap[pkgName]; ok {
		imps = append(imps, fmt.Sprintf(`"%s"`, fullPath))
	}
	// Phase001 UserClaimUnification — auth.IssueToken / RefreshToken emit a
	// `Claims: model.UserClaim{...}` literal, so the enclosing handler needs
	// the project-local model package on its import list. RefreshToken also
	// feeds claimLit into server.RefreshStore.Create via
	// buildCallRefreshCreateLines, same dependency.
	if pkgName == "auth" && (callFunc == "IssueToken" || callFunc == "RefreshToken") {
		imps = append(imps, fmt.Sprintf(`"%s/internal/model"`, g.ModulePath))
	}
	if g.IsSubscribe {
		imps = append(imps, `"fmt"`)
	} else {
		imps = append(imps, `"log/slog"`)
	}
	if g.WrapCalls {
		imps = append(imps, `"go.opentelemetry.io/otel"`)
	}
	if pkgName == "auth" && callFunc == "RefreshToken" && varName != "_" && !g.IsSubscribe && g.AccessTokenVar != "" {
		imps = append(imps, `"github.com/gin-gonic/gin"`)
	}
	return imps
}
