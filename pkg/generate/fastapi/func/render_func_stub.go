//ff:func feature=gen-fastapi type=generator control=iteration dimension=1
//ff:what RenderFuncStub — 외부 패키지 stub Python 모듈 소스 생성

package funcstub

import (
	"fmt"
	"strings"
)

// RenderFuncStub produces a Python stub module for an external package.
// Each method is generated as an async stub that raises NotImplementedError.
func RenderFuncStub(pkg string, methods []string) string {
	var b strings.Builder
	b.WriteString("# Auto-generated stub — implement with actual business logic.\n\n")

	for _, m := range methods {
		pyFunc := snakeCase(m)
		b.WriteString(fmt.Sprintf("async def %s(*args, **kwargs):\n", pyFunc))
		b.WriteString(fmt.Sprintf("    raise NotImplementedError(\"%s.%s not implemented\")\n\n\n",
			pkg, pyFunc))
	}

	return b.String()
}
