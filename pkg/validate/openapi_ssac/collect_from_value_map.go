//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-openapi
//ff:what collectFromValueMap — 문자열 값 맵에서 'request.' 접두어 참조를 fields 집합에 추가

package openapi_ssac

import "strings"

func collectFromValueMap(fields map[string]bool, m map[string]string) {
	for _, val := range m {
		if strings.HasPrefix(val, "request.") {
			fields[strings.TrimPrefix(val, "request.")] = true
		}
	}
}
