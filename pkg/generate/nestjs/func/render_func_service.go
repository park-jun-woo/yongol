//ff:func feature=gen-nestjs type=generator control=iteration dimension=1
//ff:what RenderFuncService — 외부 패키지 stub service TypeScript 소스 생성

package funcstub

import (
	"fmt"
	"strings"
)

// RenderFuncService produces a NestJS service stub file for an external package.
// Each method is generated as a stub that throws a not-implemented error.
func RenderFuncService(pkg string, methods []string) string {
	className := strings.ToUpper(pkg[:1]) + pkg[1:]
	svcName := className + "Service"

	var b strings.Builder
	b.WriteString("import { Injectable } from '@nestjs/common';\n\n")
	b.WriteString("@Injectable()\n")
	b.WriteString(fmt.Sprintf("export class %s {\n", svcName))

	for _, m := range methods {
		// lcFirst the method name
		tsMethod := strings.ToLower(m[:1]) + m[1:]
		b.WriteString(fmt.Sprintf("  async %s(...args: any[]): Promise<any> {\n", tsMethod))
		b.WriteString(fmt.Sprintf("    throw new Error('%s.%s not implemented');\n", svcName, tsMethod))
		b.WriteString("  }\n\n")
	}

	b.WriteString("}\n")
	return b.String()
}
