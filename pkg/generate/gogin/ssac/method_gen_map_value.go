//ff:func feature=gen-gogin type=util control=sequence
//ff:what methodGen.mapValue — SSaC 입력값을 Go 표현식으로 매핑 (request. prefix 분기)

package ssac

import "strings"

// mapValue converts SSaC input value to Go code using PathParams context.
func (g *methodGen) mapValue(v string) string {
	if strings.HasPrefix(v, `"`) && strings.HasSuffix(v, `"`) {
		return v
	}
	dotIdx := strings.IndexByte(v, '.')
	if dotIdx < 0 {
		return v
	}
	source := v[:dotIdx]
	field := v[dotIdx+1:]
	if source != "request" {
		return v
	}
	return g.mapRequestValue(field)
}
