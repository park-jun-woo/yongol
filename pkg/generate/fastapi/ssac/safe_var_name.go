//ff:func feature=gen-fastapi type=util control=sequence
//ff:what safeVarName — 서비스 메서드 파라미터명과 충돌 방지 (_result 접미사 추가)

package ssac

// reservedParams are the FastAPI service function parameter names that must
// not be shadowed by @get result variables.
var reservedParams = map[string]bool{
	"params":    true,
	"body":      true,
	"user":      true,
	"session":   true,
	"payload":   true,
	"event_bus": true,
}

// safeVarName appends "_result" when the variable name collides with a
// service function parameter.
func safeVarName(name string) string {
	if reservedParams[name] {
		return name + "_result"
	}
	return name
}
