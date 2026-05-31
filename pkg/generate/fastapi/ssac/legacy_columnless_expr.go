//ff:func feature=gen-fastapi type=util control=selection
//ff:what legacyColumnlessExpr — column 없는 legacy source 를 Python 식별자로 매핑

package ssac

// legacyColumnlessExpr maps a column-less SSaC source name to its Python
// identifier (request → params, currentUser → current_user, else the source).
func legacyColumnlessExpr(source string) string {
	switch source {
	case "request":
		return "params"
	case "currentUser":
		return "current_user"
	default:
		return source
	}
}
