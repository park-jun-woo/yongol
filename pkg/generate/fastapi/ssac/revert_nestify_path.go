//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what revertNestifyPath — NestJS :param → OpenAPI {param} 경로 변환 복원

package ssac

import "strings"

// revertNestifyPath converts :param back to {param} for FastAPI compatibility.
func revertNestifyPath(path string) string {
	result := path
	for {
		idx := strings.Index(result, ":")
		if idx == -1 {
			break
		}
		end := findParamEnd(result, idx+1)
		param := result[idx+1 : end]
		result = result[:idx] + "{" + param + "}" + result[end:]
	}
	return result
}
