//ff:func feature=gen-nestjs type=util control=selection
//ff:what legacyColumnlessExpr — column 없는 legacy source 를 TypeScript 식별자로 매핑

package ssac

// legacyColumnlessExpr maps a column-less SSaC source name to its TypeScript
// identifier (request → params, currentUser → user, else the source).
func legacyColumnlessExpr(source string) string {
	switch source {
	case "request":
		return "params"
	case "currentUser":
		return "user"
	default:
		return source
	}
}
