//ff:func feature=external type=generator control=sequence
//ff:what 경로 파라미터를 포함한 경로 표현식을 생성한다
package external

import (
	"fmt"
	"strings"
)

func (m methodInfo) buildPathExpr() string {
	pathParams := m.pathParams()
	if len(pathParams) == 0 {
		return fmt.Sprintf("%q", m.Path)
	}
	// Replace {param} with %v and build fmt.Sprintf
	path := m.Path
	var args []string
	for _, p := range pathParams {
		placeholder := "{" + p.Name + "}"
		if strings.Contains(path, placeholder) {
			path = strings.Replace(path, placeholder, "%v", 1)
			args = append(args, p.Name)
		}
	}
	return fmt.Sprintf("fmt.Sprintf(%q, %s)", path, strings.Join(args, ", "))
}
