//ff:func feature=external type=generator control=selection
//ff:what 메서드의 구현 코드를 생성한다
package external

import (
	"bytes"
	"fmt"
	"strings"
)

func (m methodInfo) implementation(receiverType string) string {
	var buf bytes.Buffer

	var params []string
	params = append(params, "ctx context.Context")
	for _, p := range m.Params {
		params = append(params, fmt.Sprintf("%s %s", p.Name, p.GoType))
	}

	if m.ReturnType != "" {
		fmt.Fprintf(&buf, "func (c *%s) %s(%s) (*%s, error) {\n", receiverType, m.Name, strings.Join(params, ", "), m.ReturnType)
	} else {
		fmt.Fprintf(&buf, "func (c *%s) %s(%s) error {\n", receiverType, m.Name, strings.Join(params, ", "))
	}

	// Build path with path params
	pathExpr := m.buildPathExpr()

	// Build body for POST/PUT/PATCH
	hasBody := m.HTTPMethod == "POST" || m.HTTPMethod == "PUT" || m.HTTPMethod == "PATCH"
	bodyParams := m.bodyParams()
	bodyArg := "nil"

	if hasBody && len(bodyParams) > 0 {
		writeBodyMap(&buf, bodyParams)
		bodyArg = "body"
	}

	if m.ReturnType != "" {
		writeReturnWithResult(&buf, m.HTTPMethod, pathExpr, bodyArg, m.ReturnType)
	} else {
		fmt.Fprintf(&buf, "\treturn c.do(ctx, %q, %s, %s, nil)\n", m.HTTPMethod, pathExpr, bodyArg)
	}

	buf.WriteString("}\n")
	return buf.String()
}
