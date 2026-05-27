//ff:func feature=gen-nestjs type=generator control=sequence
//ff:what RenderMain — NestJS main.ts 부트스트랩 TypeScript 소스 생성

package boot

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// RenderMain produces the NestJS main.ts bootstrap file content. It reads
// the BootPlan to determine which initialization blocks (CORS, Helmet,
// Swagger, etc.) should be included.
func RenderMain(plan *ir.BootPlan) (string, error) {
	if plan == nil {
		return "", fmt.Errorf("RenderMain: nil plan")
	}

	var b strings.Builder

	b.WriteString("import { NestFactory } from '@nestjs/core';\n")
	b.WriteString("import { AppModule } from './app.module';\n")
	b.WriteString("import { ValidationPipe } from '@nestjs/common';\n")
	b.WriteString("\n")
	b.WriteString("async function bootstrap() {\n")
	b.WriteString("  const app = await NestFactory.create(AppModule);\n\n")
	b.WriteString("  app.useGlobalPipes(\n")
	b.WriteString("    new ValidationPipe({\n")
	b.WriteString("      whitelist: true,\n")
	b.WriteString("      forbidNonWhitelisted: true,\n")
	b.WriteString("      transform: true,\n")
	b.WriteString("    }),\n")
	b.WriteString("  );\n\n")

	if hasActiveBlock(plan, "cors") {
		renderCorsBlock(&b, plan)
	}

	b.WriteString("  const port = process.env.PORT || 3000;\n")
	b.WriteString("  await app.listen(port);\n")
	b.WriteString(fmt.Sprintf("  console.log(`%s listening on port ${port}`);\n", plan.ProjectID))
	b.WriteString("}\n")
	b.WriteString("bootstrap();\n")

	return b.String(), nil
}
