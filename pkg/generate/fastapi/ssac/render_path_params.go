//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what renderPathParams — path 파라미터 이름 목록 → Python int 타입 파라미터 문자열 목록

package ssac

import "fmt"

// renderPathParams returns one Python parameter declaration per path parameter,
// each typed as int.
func renderPathParams(pathParams []string) []string {
	out := make([]string, 0, len(pathParams))
	for _, pp := range pathParams {
		out = append(out, fmt.Sprintf("%s: int", pp))
	}
	return out
}
