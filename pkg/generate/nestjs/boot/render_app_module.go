//ff:func feature=gen-nestjs type=generator control=iteration dimension=1
//ff:what RenderAppModule — NestJS app.module.ts TypeScript 소스 생성 (인프라 모듈 포함)

package boot

import (
	"fmt"
	"strings"
)

// RenderAppModule produces the root app.module.ts content. It imports
// PrismaModule, infrastructure modules (QueueModule, AuthzModule, func
// stubs), and each feature module discovered during generation.
func RenderAppModule(featureModules, infraModules []string) (string, error) {
	var b strings.Builder

	b.WriteString("import { Module } from '@nestjs/common';\n")
	b.WriteString("import { PrismaModule } from './prisma/prisma.module';\n")

	for _, im := range infraModules {
		className := strings.ToUpper(im[:1]) + im[1:] + "Module"
		b.WriteString(fmt.Sprintf("import { %s } from './%s/%s.module';\n", className, im, im))
	}

	for _, fm := range featureModules {
		className := strings.ToUpper(fm[:1]) + fm[1:] + "Module"
		b.WriteString(fmt.Sprintf("import { %s } from './%s/%s.module';\n", className, fm, fm))
	}
	b.WriteString("\n")

	b.WriteString("@Module({\n")
	b.WriteString("  imports: [\n")
	b.WriteString("    PrismaModule,\n")
	for _, im := range infraModules {
		className := strings.ToUpper(im[:1]) + im[1:] + "Module"
		b.WriteString(fmt.Sprintf("    %s,\n", className))
	}
	for _, fm := range featureModules {
		className := strings.ToUpper(fm[:1]) + fm[1:] + "Module"
		b.WriteString(fmt.Sprintf("    %s,\n", className))
	}
	b.WriteString("  ],\n")
	b.WriteString("})\n")
	b.WriteString("export class AppModule {}\n")

	return b.String(), nil
}
