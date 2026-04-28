//ff:func feature=gen-gogin type=util control=selection
//ff:what authStoreCallExpr — build the ssac/pkg/auth.RefreshRotate/Logout call expression

package ssac

import "fmt"

// authStoreCallExpr builds the `auth.Func(ctx, nil, token[, false])` call
// expression. RefreshRotate accepts the 4-arg `detectReuseLogoutAll` form
// (hard-coded false to keep zenflow behavior); Logout keeps the 3-arg shape.
func authStoreCallExpr(pkgName, callFunc, spanCtxVar, tokenArg string) string {
	switch callFunc {
	case "RefreshRotate":
		return fmt.Sprintf("%s.%s(%s, nil, %s, false)", pkgName, callFunc, spanCtxVar, tokenArg)
	default:
		return fmt.Sprintf("%s.%s(%s, nil, %s)", pkgName, callFunc, spanCtxVar, tokenArg)
	}
}
