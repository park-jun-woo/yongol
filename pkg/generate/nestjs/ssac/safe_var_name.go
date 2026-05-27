//ff:func feature=gen-nestjs type=util control=sequence
//ff:what safeVarName — 서비스 메서드 파라미터명과 충돌 방지 (_result 접미사 추가)

package ssac

// reservedParams are the NestJS service method parameter names that must not
// be shadowed by @get result variables.
var reservedParams = map[string]bool{
	"params":  true,
	"body":    true,
	"user":    true,
	"query":   true,
	"payload": true,
}

// safeVarName appends "_result" when the variable name collides with a
// service method parameter.
func safeVarName(name string) string {
	if reservedParams[name] {
		return name + "_result"
	}
	return name
}
