//ff:func feature=gen-nestjs type=util control=selection
//ff:what mapSource — SSaC source 이름 → NestJS 파라미터명 변환

package ssac

// mapSource translates SSaC source names to NestJS parameter names.
func mapSource(source string) string {
	switch source {
	case "request":
		return "params"
	case "currentUser":
		return "user"
	default:
		return source
	}
}
