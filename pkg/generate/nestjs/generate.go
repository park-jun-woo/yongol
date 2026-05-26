//ff:func feature=gen-nestjs type=command control=sequence
//ff:what Generate — NestJS 백엔드 코드 생성 (IR → TypeScript 렌더링 파이프라인)

package nestjs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/nestjs/types"
	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Generate produces NestJS backend artifacts from a parsed Fullstack.
// Pipeline:
//  1. Build IR plans (ServicePlan, BootPlan, MiddlewarePlan, InfraPlan)
//  2. Render scaffold (package.json, tsconfig.json, nest-cli.json)
//  3. Render Prisma schema from DDL tables
//  4. For each service func, render controller + service + module
//  5. Render main.ts + app.module.ts
//  6. Write all files to dir
func Generate(fs *yongol.Fullstack, dir string) error {
	psVal := prepared.New(fs)
	ps := &psVal
	reg := types.NewRegistry()

	backendDir := filepath.Join(dir, "backend")
	srcDir := filepath.Join(backendDir, "src")
	if err := os.MkdirAll(srcDir, 0o755); err != nil {
		return fmt.Errorf("mkdir src: %w", err)
	}

	bootPlan := ir.BuildBootPlan(fs, ps, "nestjs")
	_ = ir.BuildMiddlewarePlan(fs, ps)
	_ = ir.BuildInfraPlan(fs, ps)

	plansByFeature := buildPlansByFeature(fs)
	projectID := resolveProjectID(fs)

	if err := writeScaffold(backendDir, projectID); err != nil {
		return fmt.Errorf("scaffold: %w", err)
	}
	if err := writePrismaSchema(fs, backendDir); err != nil {
		return fmt.Errorf("prisma schema: %w", err)
	}
	if err := writePrismaModule(srcDir); err != nil {
		return fmt.Errorf("prisma module: %w", err)
	}
	featureNames, err := writeFeatureModules(plansByFeature, srcDir, reg)
	if err != nil {
		return err
	}
	return writeBootFiles(srcDir, bootPlan, featureNames)
}
