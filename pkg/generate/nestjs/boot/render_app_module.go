//ff:func feature=gen-nestjs type=generator control=iteration dimension=1
//ff:what RenderAppModule — NestJS app.module.ts TypeScript 소스 생성

package boot

import (
	"fmt"
	"strings"
)

// RenderAppModule produces the root app.module.ts content. It imports
// PrismaModule and each feature module discovered during generation.
func RenderAppModule(featureModules []string) (string, error) {
	var b strings.Builder

	b.WriteString("import { Module } from '@nestjs/common';\n")
	b.WriteString("import { PrismaModule } from './prisma/prisma.module';\n")

	for _, fm := range featureModules {
		className := strings.ToUpper(fm[:1]) + fm[1:] + "Module"
		b.WriteString(fmt.Sprintf("import { %s } from './%s/%s.module';\n", className, fm, fm))
	}
	b.WriteString("\n")

	b.WriteString("@Module({\n")
	b.WriteString("  imports: [\n")
	b.WriteString("    PrismaModule,\n")
	for _, fm := range featureModules {
		className := strings.ToUpper(fm[:1]) + fm[1:] + "Module"
		b.WriteString(fmt.Sprintf("    %s,\n", className))
	}
	b.WriteString("  ],\n")
	b.WriteString("})\n")
	b.WriteString("export class AppModule {}\n")

	return b.String(), nil
}
