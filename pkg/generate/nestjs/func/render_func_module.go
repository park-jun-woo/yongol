//ff:func feature=gen-nestjs type=generator control=sequence
//ff:what RenderFuncModule — 외부 패키지 stub module TypeScript 소스 생성

package funcstub

import (
	"fmt"
	"strings"
)

// RenderFuncModule produces a NestJS module file for an external package stub.
func RenderFuncModule(pkg string, methods []string) string {
	className := strings.ToUpper(pkg[:1]) + pkg[1:]
	svcName := className + "Service"
	modName := className + "Module"

	var b strings.Builder
	b.WriteString("import { Global, Module } from '@nestjs/common';\n")
	b.WriteString(fmt.Sprintf("import { %s } from './%s.service';\n\n", svcName, pkg))
	b.WriteString("@Global()\n")
	b.WriteString("@Module({\n")
	b.WriteString(fmt.Sprintf("  providers: [%s],\n", svcName))
	b.WriteString(fmt.Sprintf("  exports: [%s],\n", svcName))
	b.WriteString("})\n")
	b.WriteString(fmt.Sprintf("export class %s {}\n", modName))
	return b.String()
}
