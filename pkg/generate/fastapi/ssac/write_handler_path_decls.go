//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what writeHandlerPathDecls — handler 함수의 path 파라미터 선언 출력

package ssac

import (
	"fmt"
	"strings"
)

// writeHandlerPathDecls writes one typed path parameter declaration per name.
func writeHandlerPathDecls(b *strings.Builder, pathParams []string) {
	for _, pp := range pathParams {
		b.WriteString(fmt.Sprintf("    %s: int,\n", pp))
	}
}
