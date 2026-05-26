//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what findParamEnd — 경로 문자열에서 파라미터 이름 끝 위치 탐색

package ssac

// findParamEnd finds the end index of a path parameter name starting at
// start. The parameter ends at the next '/' or end of string.
func findParamEnd(path string, start int) int {
	end := start
	for end < len(path) && path[end] != '/' {
		end++
	}
	return end
}
