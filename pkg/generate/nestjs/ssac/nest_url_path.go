//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what nestURLPath — OpenAPI {param} → NestJS :param 경로 변환

package ssac

import "strings"

// nestURLPath converts OpenAPI {param} notation to NestJS :param.
func nestURLPath(path string) string {
	result := path
	for {
		start := strings.Index(result, "{")
		if start == -1 {
			break
		}
		end := strings.Index(result[start:], "}")
		if end == -1 {
			break
		}
		param := result[start+1 : start+end]
		result = result[:start] + ":" + param + result[start+end+1:]
	}
	return result
}
