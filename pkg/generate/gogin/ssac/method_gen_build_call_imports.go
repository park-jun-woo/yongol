//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildCallImports — @call emission 시 필요한 import 경로 수집

package ssac

import "fmt"

func (g *methodGen) buildCallImports(pkgName, callFunc, varName string) []string {
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
		imps = append(imps, `"go.opentelemetry.io/otel"`)
	}
	if pkgName == "auth" && callFunc == "RefreshToken" && varName != "_" && !g.IsSubscribe && g.AccessTokenVar != "" {
		imps = append(imps, `"github.com/gin-gonic/gin"`)
	}
	return imps
}
