//ff:func feature=external type=generator control=iteration dimension=1
//ff:what 요청 바디 맵 코드를 버퍼에 작성한다
package external

import (
	"bytes"
	"fmt"
)

func writeBodyMap(buf *bytes.Buffer, bodyParams []paramInfo) {
	buf.WriteString("\tbody := map[string]any{")
	for i, p := range bodyParams {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "%q: %s", p.Name, p.Name)
	}
	buf.WriteString("}\n")
}
