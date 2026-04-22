//ff:func feature=external type=generator control=sequence
//ff:what 반환값이 있는 메서드의 do 호출 및 반환 코드를 작성한다
package external

import (
	"bytes"
	"fmt"
)

func writeReturnWithResult(buf *bytes.Buffer, httpMethod, pathExpr, bodyArg, returnType string) {
	fmt.Fprintf(buf, "\tvar resp %s\n", returnType)
	fmt.Fprintf(buf, "\tif err := c.do(ctx, %q, %s, %s, &resp); err != nil {\n", httpMethod, pathExpr, bodyArg)
	buf.WriteString("\t\treturn nil, err\n")
	buf.WriteString("\t}\n")
	buf.WriteString("\treturn &resp, nil\n")
}
