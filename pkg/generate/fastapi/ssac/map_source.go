//ff:func feature=gen-fastapi type=util control=selection
//ff:what mapSource — SSaC source 이름 → FastAPI 파라미터명 변환

package ssac

// mapSource translates SSaC source names to FastAPI parameter names.
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
