//ff:func feature=gen-gogin type=util control=sequence
//ff:what authStoreImports — import list for the emitted auth RefreshRotate/Logout call block

package ssac

import "fmt"

// authStoreImports collects the imports needed by the auth RefreshRotate /
// Logout call block. Mirrors the original sequential collection in
// buildAuthRefreshStoreCall so output order stays stable.
func (g *methodGen) authStoreImports(pkgName, callFunc string) []string {
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
	return imps
}
