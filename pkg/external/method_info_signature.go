//ff:func feature=external type=generator control=sequence
//ff:what 메서드의 인터페이스 시그니처 문자열을 생성한다
package external

import (
	"fmt"
	"strings"
)

func (m methodInfo) signature() string {
	var params []string
	params = append(params, "ctx context.Context")
	for _, p := range m.Params {
		params = append(params, fmt.Sprintf("%s %s", p.Name, p.GoType))
	}
	if m.ReturnType != "" {
		return fmt.Sprintf("%s(%s) (*%s, error)", m.Name, strings.Join(params, ", "), m.ReturnType)
	}
	return fmt.Sprintf("%s(%s) error", m.Name, strings.Join(params, ", "))
}
